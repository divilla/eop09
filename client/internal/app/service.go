package app

import (
	"encoding/json"
	"fmt"
	i "github.com/divilla/eop09/client/internal/interfaces"
	pb "github.com/divilla/eop09/entityproto"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"golang.org/x/net/context"
	"io"
	"net/http"
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

func (s *service) index(ctx context.Context, currentPage, perPage int64) ([]byte, *pb.IndexResponse, error) {
	var key string
	var value json.RawMessage
	var err error

	indexResponse, err := s.client.Index(ctx, &pb.IndexRequest{
		CurrentPage: currentPage,
		PerPage:     perPage,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to receive IndexResponse: %w", err)
	}

	res := []byte(`{}`)
	for _, v := range indexResponse.GetResults() {
		key, value, err = decodeEntityJson(v.GetJson())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to decode Entity json: %w", err)
		}

		res, err = sjson.SetRawBytes(res, key, value)
		if err != nil {
			return nil, nil, fmt.Errorf("error building index json response: %w", err)
		}
	}

	return res, indexResponse, nil
}

func (s *service) get(ctx context.Context, key string) ([]byte, error) {
	entity, err := s.client.Get(ctx, &pb.KeyRequest{Key: key})
	if err != nil {
		return nil, err
	}

	key, value, err := decodeEntityJson(entity.GetJson())
	if err != nil {
		return nil, fmt.Errorf("failed to decode Entity json: %w", err)
	}

	return sjson.SetRawBytes([]byte(`{}`), key, value)
}

func (s *service) create(ctx context.Context, result *gjson.Result) (*pb.CommandResponse, error) {
	res, err := encodeEntityJson(result)
	if err != nil {
		return nil, err
	}

	return s.client.Create(ctx, &pb.Entity{Json: res})
}

func (s *service) patch(ctx context.Context, currentKey string, result *gjson.Result) (*pb.CommandResponse, error) {
	res, err := encodeEntityJson(result)
	if err != nil {
		return nil, err
	}

	return s.client.Patch(ctx, &pb.KeyEntityRequest{
		Key:  currentKey,
		Json: res,
	})
}

func (s *service) put(ctx context.Context, currentKey string, result *gjson.Result) (*pb.CommandResponse, error) {
	res, err := encodeEntityJson(result)
	if err != nil {
		return nil, err
	}

	return s.client.Put(ctx, &pb.KeyEntityRequest{
		Key:  currentKey,
		Json: res,
	})
}

func (s *service) delete(ctx context.Context, key string) (*pb.CommandResponse, error) {
	return s.client.Delete(ctx, &pb.KeyRequest{
		Key: key,
	})
}

func (s *service) importer(ctx context.Context) (json.RawMessage, bool, error) {
	var index uint64
	var key string
	var value json.RawMessage
	var err error

	if err = s.client.Ping(); err != nil {
		return nil, false, echo.NewHTTPError(http.StatusGone, err)
	}

	impCli, err := s.client.Import(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("failed to open grpc upstream: %w", err)
	}

	err = s.reader.Start()
	if err != nil {
		return nil, false, fmt.Errorf("json reader failed to start: %w", err)
	}
	for {
		err = s.reader.Read(&index, &key, &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, false, fmt.Errorf("json file read error: %w", err)
		}

		value, err = encodeEntityKeyValue(key, &value)
		if err != nil {
			return nil, false, fmt.Errorf("import json encoding failed: %w", err)
		}

		err = impCli.Send(&pb.Entity{Json: value})
		if err != nil {
			return nil, false, fmt.Errorf("gRPC upstream send failed: %w", err)
		}
	}
	err = s.reader.Close()
	if err != nil {
		return nil, false, fmt.Errorf("json reader failed to close: %w", err)
	}

	ir, err := impCli.CloseAndRecv()
	if err != nil {
		return nil, false, fmt.Errorf("gRPC import client failed to close and receive: %w", err)
	}

	res, success, err := parseImportResponse(ir)
	if err != nil {
		return nil, false, fmt.Errorf("failed to parse import result: %w", err)
	}

	return res, success, nil
}
