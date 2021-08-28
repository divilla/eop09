package importer

import (
	"encoding/json"
	"fmt"
	i "github.com/divilla/eop09/client/internal/interfaces"
	"github.com/divilla/eop09/client/pkg/jsondecimals"
	"github.com/divilla/eop09/crudproto"
	"github.com/tidwall/sjson"
	"golang.org/x/net/context"
	"io"
)

type service struct {
	client i.GRPCClient
	reader i.JsonReader
	logger i.Logger
}

func newService(client i.GRPCClient, reader i.JsonReader, logger i.Logger) *service {
	return &service{
		client: client,
		reader: reader,
		logger: logger,
	}
}

func (s *service) list(ctx context.Context, pageNumber, pageSize int64) ([]byte, error) {
	req, err := s.client.List(ctx, &crudproto.ListRequest{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	})
	if err != nil {
		return nil, err
	}

	res := []byte(`{}`)
	for _, v := range req.GetResults() {
		value, err := jsondecimals.Unquote(v.GetValue(), "coordinates")
		if err != nil {
			s.logger.Error(err)
		}

		res, err = sjson.SetRawBytes(res, v.GetKey(), value)
		if err != nil {
			s.logger.Errorf("error setting json key & value: %w", err)
		}
	}

	return res, nil
}

func (s *service) importer(ctx context.Context) (*crudproto.CommandResponse, error) {
	var index uint64
	var key string
	var value json.RawMessage
	var err error

	impCli, err := s.client.Import(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to open grpc upstream: %w", err)
	}

	s.reader.Reset()
	for {
		err = s.reader.Read(&index, &key, &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			s.logger.Errorf("json file read error: %w", err)
		}

		value, err = jsondecimals.Quote(value, "coordinates")
		if err != nil {
			s.logger.Error(err)
		}

		err = impCli.Send(&crudproto.Entity{
			Key: key,
			Value: value,
		})
		if err != nil {
			s.logger.Errorf("grpc upstream send failed: %w", err)
		}
	}

	res, err := impCli.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return res, nil
}
