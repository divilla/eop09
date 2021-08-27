package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/divilla/eop09/crudproto"
	"github.com/divilla/eop09/server/internal/domain"
	"github.com/divilla/eop09/server/internal/interfaces"
	"github.com/divilla/eop09/server/pkg/cmongo"
	"io"
)

type Server struct {
	pb.UnimplementedRPCServer
	repository cmongo.Repository
	logger     interfaces.Logger
}

func NewServer(repository cmongo.Repository, logger interfaces.Logger) *Server {
	return &Server{
		repository: repository,
		logger:     logger,
	}
}

//List returns batch of Entities
func (s *Server) List(ctx context.Context, listRequest *pb.ListRequest) (*pb.ListResponse, error) {
	var es []domain.Port

	err := s.repository.List(ctx, listRequest.PageNumber, listRequest.PageSize, &es)
	if err != nil {
		return nil, err
	}
	
	lr := &pb.ListResponse{
		Results: make([]*pb.Entity, len(es)),
	}
	for k, v := range es {
		e, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		lr.Results[k] = &pb.Entity{
			Key:   v.Id,
			Value: e,
		}
	}

	return lr, nil
}

//Import implements RPC_ImportServer
func (s *Server) Import(stream pb.RPC_ImportServer) error {
	var port *domain.Port
	var errs []string
	i := uint64(0)

	for {
		entity, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CommandResponse{
				Success:  true,
				RowsAffected: i,
				Errors: errs,
			})
		}
		if err != nil {
			s.logger.Errorf("Error while receiving stream: ", err)
			return stream.SendAndClose(&pb.CommandResponse{
				Success:  false,
				RowsAffected: i,
				Errors: []string{err.Error()},
			})
		}

		port, err = entityToPort(entity)
		if err != nil {
			err = fmt.Errorf("failed to unmarshall domain.Port '%s' with error: %w", string(entity.GetValue()), err)
			s.logger.Error(err)
			errs = append(errs, err.Error())
		}

		err = s.repository.UpsertOne(context.TODO(), entity.Key, port)
		if err != nil {
			err = fmt.Errorf("failed to save domain.Port '%s' with error: %w", string(entity.GetValue()), err)
			s.logger.Error(err)
			errs = append(errs, err.Error())
		}

		i++
	}
}

func entityToPort(e *pb.Entity) (*domain.Port, error) {
	p := new(domain.Port)
	err := json.Unmarshal(e.GetValue(), p)
	if err != nil {
		return nil, err
	}

	p.Id = e.Key
	return p, nil
}
