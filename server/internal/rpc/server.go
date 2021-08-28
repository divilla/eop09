package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/divilla/eop09/entityproto"
	"github.com/divilla/eop09/server/internal/domain"
	"github.com/divilla/eop09/server/internal/dto"
	i "github.com/divilla/eop09/server/internal/interfaces"
	"github.com/tidwall/sjson"
	"io"
	"math"
	"net/http"
	"strings"
)

type Server struct {
	pb.UnimplementedRPCServer
	repository i.Repository
	logger     i.Logger
}

//NewServer creates new gRPC server with repository and logger
func NewServer(repository i.Repository, logger i.Logger) *Server {
	return &Server{
		repository: repository,
		logger:     logger,
	}
}

//Index returns batch of Entities
func (s *Server) Index(ctx context.Context, in *pb.IndexRequest) (*pb.IndexResponse, error) {
	var es []domain.Port

	totalResults, err := s.repository.CountAll(ctx)
	if err != nil {
		return nil, err
	}
	totalPages := int64(math.Ceil(float64(totalResults) / float64(in.PageSize)))

	err = s.repository.List(ctx, in.Page, in.PageSize, &es)
	if err != nil {
		return nil, err
	}

	lr := &pb.IndexResponse{
		Page:         in.Page,
		PageSize:     in.PageSize,
		Results:      make([]*pb.Entity, len(es)),
		TotalResults: totalResults,
		TotalPages:   totalPages,
	}

	for k, v := range es {
		e, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		lr.Results[k] = &pb.Entity{
			Key:   v.Key,
			Value: e,
		}
	}

	return lr, nil
}

//Get returns single entity found by key (id)
func (s *Server) Get(ctx context.Context, in *pb.KeyRequest) (*pb.Entity, error) {
	var port domain.Port
	err := s.repository.FindOne(ctx, in.GetKey(), &port)
	if err != nil {
		return nil, err
	}

	value, err := json.Marshal(port)
	if err != nil {
		return nil, err
	}

	return &pb.Entity{
		Key:   port.Key,
		Value: value,
	}, nil
}

//Create creates new document in db
func (s *Server) Create(ctx context.Context, in *pb.Entity) (*pb.CommandResponse, error) {
	port := new(dto.PortDto)
	err := unmarshalAndValidatePortDto(port, in, true)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateOne(ctx, port)
	if err != nil && strings.Contains(err.Error(), "E11000") {
		return nil, dto.NewJsonError(http.StatusBadRequest, "document with requested key already exists")
	}
	if err != nil {
		return nil, err
	}

	return newCommandResponse(1), nil
}

//Patch updates values of existing document with the same id
func (s *Server) Patch(ctx context.Context, in *pb.KeyEntity) (*pb.CommandResponse, error) {
	port := new(dto.PortDto)
	err := s.repository.FindOne(ctx, in.GetOldKey(), port)
	if err != nil {
		return nil, err
	}

	err = unmarshalAndValidateKeyPortDto(port, in, true)
	if err != nil {
		return nil, err
	}

	err = s.repository.ReplaceOne(ctx, in.GetOldKey(), port)
	if err != nil {
		err = fmt.Errorf("failed to update value in db: %w", err)
		return nil, err
	}

	return newCommandResponse(1), nil
}

//Put replaces document with the new one with the same id
func (s *Server) Put(ctx context.Context, in *pb.KeyEntity) (*pb.CommandResponse, error) {
	port := new(dto.PortDto)
	err := unmarshalAndValidateKeyPortDto(port, in, true)
	if err != nil {
		return nil, err
	}

	err = s.repository.ReplaceOne(ctx, in.GetOldKey(), port)
	if err != nil {
		err = fmt.Errorf("failed to update value in db: %w", err)
		return nil, err
	}

	return newCommandResponse(1), nil
}

//Delete deletes entity found by key (id)
func (s *Server) Delete(ctx context.Context, in *pb.KeyRequest) (*pb.CommandResponse, error) {
	err := s.repository.DeleteOne(ctx, in.GetKey())
	if err != nil {
		err = fmt.Errorf("failed to delete key: %w", err)
		return nil, err
	}

	return newCommandResponse(1), nil
}

//Import implements RPC_ImportServer
func (s *Server) Import(stream pb.RPC_ImportServer) error {
	res := &pb.ImportResponse{
		Success:      true,
		RowsAffected: int64(0),
		Errors:       "",
	}
	jErrors := newJsonErrors()

	for {
		entity, err := stream.Recv()
		if err == io.EOF {
			res.Errors = jErrors.Errors()
			return stream.SendAndClose(res)
		}
		if err != nil {
			err = fmt.Errorf("error while receiving stream: %w", err)
			s.logger.Error(err)
			return err
		}

		port := new(dto.PortDto)
		err = unmarshalAndValidatePortDto(port, entity, false)
		if err != nil {
			s.logger.Error(err)
			res.Success = false
			jErrors.Add(entity.Key, entity.Value)
		}

		err = s.repository.UpsertOne(context.TODO(), entity.Key, port)
		if err != nil {
			err = fmt.Errorf("failed to save domain.Port '%s' with error: %w", string(entity.GetValue()), err)
			s.logger.Error(err)
			res.Success = false
			jErrors.Add(entity.Key, entity.Value)
		} else {
			res.RowsAffected++
		}
	}
}

func unmarshalAndValidateKeyPortDto(p *dto.PortDto, e *pb.KeyEntity, validate bool) error {
	return unmarshalAndValidatePortDto(p, &pb.Entity{
		Key:   e.GetKey(),
		Value: e.GetValue(),
	}, validate)
}

func unmarshalAndValidatePortDto(p *dto.PortDto, e *pb.Entity, validate bool) error {
	err := json.Unmarshal(e.GetValue(), p)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal domain.Port: %w", err)
		return err
	}
	p.Key = e.Key

	if validate {
		validationErrors := p.Validate()
		if validationErrors == nil {
			return nil
		}

		jsonErrors, err := validationErrors.MarshalJSON()
		if err != nil {
			return fmt.Errorf("marshaling validation errors failed: %w", err)
		}

		return dto.NewValidationErrors(jsonErrors)
	}

	return nil
}

func newCommandResponse(rows int64) *pb.CommandResponse {
	return &pb.CommandResponse{
		RowsAffected: rows,
	}
}

type jsonErrors json.RawMessage

func newJsonErrors() *jsonErrors {
	return &jsonErrors{}
}

func (j *jsonErrors) Add(key string, value []byte) *jsonErrors {
	*j, _ = sjson.SetRawBytes(*j, key, value)
	return j
}

func (j *jsonErrors) Errors() string {
	if len(*j) == 0 {
		return ""
	}
	return string(*j)
}
