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
	"io"
	"net/http"
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

	e.GET("/import", ctrl.importer)
}

func (c *controller) importer(cc echo.Context) error {
	ctx := cc.(*cmiddleware.Context)
	var index uint64
	var key string
	var value json.RawMessage
	var err error

	impCli, err := c.client.ImportClient(ctx.RequestContext())
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

		err = impCli.Send(&crudproto.Entity{Result: value})
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
