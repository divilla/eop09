package rpc

import (
	pb "github.com/divilla/eop09/crudproto"
	"github.com/divilla/eop09/server/internal/domain"
	"github.com/divilla/eop09/server/internal/interfaces"
	"io"
)

type Server struct {
	pb.UnimplementedRPCServer
	logger interfaces.Logger
}

func NewServer(logger interfaces.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

//Import implements RPC_ImportServer
func (s *Server) Import(stream pb.RPC_ImportServer) error {
	var port domain.Port
	i := uint64(0)
	errs := ""

	for {
		entity, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Success:  true,
				AffectedRows: i,
				Error: errs,
			})
		}
		if err != nil {
			s.logger.Errorf("Error while receiving stream: ", err)
			return stream.SendAndClose(&pb.Response{
				Success:  false,
				AffectedRows: i,
				Error: "Error while receiving stream: " + err.Error(),
			})
		}

		if err = port.Unmarshal(entity.GetResult()); err != nil {
			s.logger.Errorf("failed to unmarshall Port: %s", string(entity.GetResult()))
		}

		i++
	}
}
