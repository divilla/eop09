package app

import (
	"encoding/json"
	"fmt"
	i "github.com/divilla/eop09/client/internal/interfaces"
	"github.com/divilla/eop09/client/pkg/jsondecimals"
	"github.com/divilla/eop09/crudproto"
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

func (s *service) index(ctx context.Context, page, pageSize int64) ([]byte, *entityproto.IndexResponse, error) {
	var value json.RawMessage
	var err error

	res, err := s.client.Index(ctx, &entityproto.IndexRequest{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, nil, err
	}

	response := []byte(`{}`)
	for _, v := range res.GetResults() {
		value, err = jsondecimals.Unquote(v.GetValue(), "coordinates")
		if err != nil {
			s.logger.Error(err)
		}

		response, err = sjson.SetRawBytes(response, v.GetKey(), value)
		if err != nil {
			s.logger.Errorf("error setting json key & value: %w", err)
		}
	}

	return response, res, nil
}

func (s *service) get(ctx context.Context, key string) ([]byte, error) {
	entity, err := s.client.Get(ctx, &entityproto.KeyRequest{Key: key})
	if err != nil {
		return nil, err
	}

	value, err := jsondecimals.Unquote(entity.GetValue(), "coordinates")
	if err != nil {
		return nil, err
	}

	res := []byte(`{}`)
	return sjson.SetRawBytes(res, entity.GetKey(), value)
}

func (s *service) create(ctx context.Context, result gjson.Result) (*entityproto.CommandResponse, error) {
	var key, value string
	result.ForEach(func(k, v gjson.Result) bool {
		key = k.String()
		value = v.Raw
		return false
	})

	val, err := jsondecimals.Quote([]byte(value), "coordinates")
	if err != nil {
		return nil, err
	}

	return s.client.Create(ctx, &entityproto.Entity{
		Key:   key,
		Value: val,
	})
}

func (s *service) patch(ctx context.Context, oldKey string, result gjson.Result) (*entityproto.CommandResponse, error) {
	var key, value string
	result.ForEach(func(k, v gjson.Result) bool {
		key = k.String()
		value = v.Raw
		return false
	})

	val, err := jsondecimals.Quote([]byte(value), "coordinates")
	if err != nil {
		return nil, err
	}

	return s.client.Patch(ctx, &entityproto.KeyEntity{
		OldKey: oldKey,
		Key:    key,
		Value:  val,
	})
}

func (s *service) put(ctx context.Context, oldKey string, result gjson.Result) (*entityproto.CommandResponse, error) {
	var key, value string
	result.ForEach(func(k, v gjson.Result) bool {
		key = k.String()
		value = v.Raw
		return false
	})

	val, err := jsondecimals.Quote([]byte(value), "coordinates")
	if err != nil {
		return nil, err
	}

	return s.client.Put(ctx, &entityproto.KeyEntity{
		OldKey: oldKey,
		Key:    key,
		Value:  val,
	})
}

func (s *service) delete(ctx context.Context, key string) (*entityproto.CommandResponse, error) {
	return s.client.Delete(ctx, &entityproto.KeyRequest{
		Key: key,
	})
}

func (s *service) importer(ctx context.Context) (*entityproto.ImportResponse, error) {
	var index uint64
	var key string
	var value json.RawMessage
	var err error

	if !s.client.IsConnected() {
		return nil, echo.NewHTTPError(http.StatusGone, "gRPC client not connected, please try again later")
	}

	impCli, err := s.client.Import(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to open grpc upstream: %w", err)
	}

	err = s.reader.Start()
	if err != nil {
		return nil, fmt.Errorf("json reader failed to start: %w", err)
	}
	for {
		err = s.reader.Read(&index, &key, &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("json file read error: %w", err)
		}

		value, err = jsondecimals.Quote(value, "coordinates")
		if err != nil {
			s.logger.Error(err)
		}

		err = impCli.Send(&entityproto.Entity{
			Key:   key,
			Value: value,
		})
		if err != nil {
			s.logger.Errorf("grpc upstream send failed: %w", err)
		}
	}
	err = s.reader.Close()
	if err != nil {
		return nil, fmt.Errorf("json reader failed to close: %w", err)
	}

	res, err := impCli.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return res, nil
}
