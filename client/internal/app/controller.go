package importer

import (
	i "github.com/divilla/eop09/client/internal/interfaces"
	. "github.com/divilla/eop09/client/pkg/cecho"
	"github.com/labstack/echo/v4"
	"net/http"
)

type controller struct {
	service *service
	logger i.Logger
}

func Controller(e *echo.Echo, client i.GRPCClient, reader i.JsonReader) {
	ctrl := &controller{
		service: newService(client, reader, e.Logger),
		logger: e.Logger,
	}

	g := e.Group("/ports")
	g.GET("", H(ctrl.list))
	g.GET("/import", H(ctrl.importer))
}

func (c *controller) list(ctx i.Context) error {
	res, err := c.service.list(ctx.RequestContext(),
		ctx.QueryParamInt64("page", 1),
		ctx.QueryParamInt64("results", 30))
	if err != nil {
		return err
	}

	return ctx.JSONBytes(http.StatusOK, res)
}

func (c *controller) importer(ctx i.Context) error {
	res, err := c.service.importer(ctx.RequestContext())
	if err != nil {
		return err
	}

	status := http.StatusOK
	if !res.Success {
		status = http.StatusBadRequest
	}

	return ctx.JSON(status, res)
}
