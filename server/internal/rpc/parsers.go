package rpc

import (
	"encoding/json"
	"fmt"
	pb "github.com/divilla/eop09/entityproto"
	"github.com/divilla/eop09/server/internal/dto"
)

func unmarshalAndValidateKeyEntity(p *dto.Port, e *pb.KeyEntityRequest) error {
	return unmarshalAndValidateEntity(p, &pb.Entity{Json: e.Json})
}

func unmarshalAndValidateEntity(p *dto.Port, e *pb.Entity) error {
	err := json.Unmarshal(e.GetJson(), p)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal domain.Port: %w", err)
		return err
	}

	errors, err := p.ValidateAndMarshal()
	if err != nil {
		panic(err)
	}

	if errors == nil {
		return nil
	}

	return dto.NewValidationErrors(errors)
}

func newCommandResponse(rows int64) *pb.CommandResponse {
	return &pb.CommandResponse{
		RowsAffected: rows,
	}
}

