package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/divilla/eop09/entityproto"
	"github.com/divilla/eop09/server/internal/dto"
	i "github.com/divilla/eop09/server/internal/interfaces"
	"github.com/tidwall/gjson"
	"io"
	"math"
	"time"
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
	var es []dto.Port

	totalCount, err := s.repository.CountAll(ctx)
	if err != nil {
		return nil, err
	}
	pageCount := int64(math.Ceil(float64(totalCount) / float64(in.GetPerPage())))

	err = s.repository.List(ctx, in.GetCurrentPage(), in.GetPerPage(), &es)
	if err != nil {
		return nil, err
	}

	lr := &pb.IndexResponse{
		Results:     make([]*pb.Entity, len(es)),
		CurrentPage: in.CurrentPage,
		PerPage:     in.PerPage,
		TotalCount:  totalCount,
		PageCount:   pageCount,
	}

	for k, v := range es {
		e, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		lr.Results[k] = &pb.Entity{Json: e}
	}

	return lr, nil
}

//Get returns single entity found by key (id)
func (s *Server) Get(ctx context.Context, in *pb.KeyRequest) (*pb.Entity, error) {
	var port dto.Port
	err := s.repository.FindOne(ctx, in.GetKey(), &port)
	if err != nil {
		return nil, err
	}

	value, err := json.Marshal(port)
	if err != nil {
		return nil, err
	}

	return &pb.Entity{Json: value}, nil
}

//Create creates new document in db
func (s *Server) Create(ctx context.Context, in *pb.Entity) (*pb.CommandResponse, error) {
	port := new(dto.Port)
	err := unmarshalAndValidateEntity(port, in)
	if err != nil {
		return nil, err
	}

	err = s.repository.CreateOne(ctx, port)
	if err != nil {
		return nil, err
	}

	return newCommandResponse(1), nil
}

//Patch updates values of existing document with the same id
func (s *Server) Patch(ctx context.Context, in *pb.KeyEntityRequest) (*pb.CommandResponse, error) {
	port := new(dto.Port)
	err := s.repository.FindOne(ctx, in.GetKey(), port)
	if err != nil {
		return nil, err
	}

	err = unmarshalAndValidateKeyEntity(port, in)
	if err != nil {
		return nil, err
	}

	err = s.repository.ReplaceOne(ctx, in.GetKey(), port)
	if err != nil {
		err = fmt.Errorf("failed to update value in db: %w", err)
		return nil, err
	}

	return newCommandResponse(1), nil
}

//Put replaces document with the new one with the same id
func (s *Server) Put(ctx context.Context, in *pb.KeyEntityRequest) (*pb.CommandResponse, error) {
	port := new(dto.Port)
	err := unmarshalAndValidateKeyEntity(port, in)
	if err != nil {
		return nil, err
	}

	err = s.repository.ReplaceOne(ctx, in.GetKey(), port)
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
	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
	defer cancel()

	res := &pb.ImportResponse{
		Success:      true,
		RowsAffected: int64(0),
	}
	jErrors := newJsonErrors()

	for {
		entity, err := stream.Recv()
		if err == io.EOF {
			res.Errors = jErrors.Errors()
			return stream.SendAndClose(res)
		}
		if err != nil {
			return fmt.Errorf("error receiving import stream: %w", err)
		}

		port := new(dto.Port)
		if err = unmarshalAndValidateEntity(port, entity); err != nil {
			s.logger.Error(err)
			res.Success = false
			if err := jErrors.Add(entity.GetJson()); err != nil {
				return err
			}
			continue
		}

		err = s.repository.UpsertOne(ctx, gjson.GetBytes(entity.GetJson(), KeyPath).String(), port)
		if err != nil {
			s.logger.Error(err)
			res.Success = false
			if err := jErrors.Add(entity.GetJson()); err != nil {
				return err
			}
		} else {
			res.RowsAffected++
		}
	}
}
