package importer

import (
	"encoding/json"
	"fmt"
	"github.com/divilla/eop09/client/config"
	"github.com/divilla/eop09/client/internal/grpcc"
	"github.com/divilla/eop09/client/internal/interfaces"
	"github.com/divilla/eop09/client/pkg/cmiddleware"
	"github.com/divilla/eop09/client/pkg/largejsonreader"
	"github.com/divilla/eop09/crudproto"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type controller struct {
	client *grpcc.Client
	logger interfaces.Logger
}

func Controller(e *echo.Echo, client *grpcc.Client) {
	ctrl := &controller{
		client: client,
		logger: e.Logger,
		//ser: &service{},
	}

	e.GET("/list", ctrl.list)
	e.GET("/list/:pageNumber", ctrl.list)
	e.GET("/list/:pageNumber/:pageSize", ctrl.list)
	e.GET("/import", ctrl.importer)
}

func (c *controller) list(cc echo.Context) error {
	ctx := cc.(*cmiddleware.Context)

	pageNumber, err := strconv.ParseInt(ctx.Param("pageNumber"), 10, 64)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}

	pageSize, err := strconv.ParseInt(ctx.Param("pageSize"), 10, 64)
	if err != nil || pageSize < 1 {
		pageSize = 30
	}

	lr, err := c.client.List(ctx.RequestContext(), &crudproto.ListRequest{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	})
	if err != nil {
		return err
	}

	j := []byte(`{}`)
	for _, v := range lr.GetResults() {
		value := v.GetValue()

		var cords []string
		rawCords := "[]"
		gjson.GetBytes(value, "coordinates").ForEach(func(key, val gjson.Result) bool {
			cords = append(cords, val.String())
			return true
		})
		if len(cords) > 0 {
			rawCords = "[" + strings.Join(cords, ",") + "]"
		}
		value, _ = sjson.SetRawBytes(value, "coordinates", []byte(rawCords))

		j, err = sjson.SetRawBytes(j, v.GetKey(), value)
		if err != nil {
			c.logger.Errorf("error setting json value: %w", err)
		}
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	_, err = ctx.Response().Write(j)
	return err
}

func (c *controller) importer(cc echo.Context) error {
	ctx := cc.(*cmiddleware.Context)
	var index uint64
	var key string
	var value json.RawMessage
	var err error

	impCli, err := c.client.Import(ctx.RequestContext())
	if err != nil {
		return fmt.Errorf("failed to open grpc upstream: %w", err)
	}

	jReader, err := largejsonreader.Read(config.App.JsonDataFile)
	if err != nil {
		return fmt.Errorf("failed to start largeJsonFile reader: %w", err)
	}

	for {
		err = jReader.Next(&index, &key, &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			c.logger.Errorf("json file read error: %w", err)
		}

		var cords []string
		gjson.GetBytes(value, "coordinates").ForEach(func(key, val gjson.Result) bool {
			cords = append(cords, val.Raw)
			return true
		})
		value, err = sjson.SetBytes(value, "coordinates", cords)
		if err != nil {
			c.logger.Errorf("unable to set coordinates: %w", err)
		}

		err = impCli.Send(&crudproto.Entity{
			Key: key,
			Value: value,
		})
		if err != nil {
			c.logger.Errorf("grpc upstream send failed: %w", err)
		}
	}

	if err = jReader.Close(); err != nil {
		panic(err)
	}

	res, err := impCli.CloseAndRecv()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}
