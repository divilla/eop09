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

	e.GET("/list", H(ctrl.list))
	e.GET("/list/:pageNumber", H(ctrl.list))
	e.GET("/list/:pageNumber/:pageSize", H(ctrl.list))
	e.GET("/import", H(ctrl.importer))
}

func (c *controller) list(ctx Context) error {
	res, err := c.service.list(ctx.RequestContext(),
		ctx.ParamInt64("pageNumber", 1),
		ctx.ParamInt64("pageNumber", 30))
	if err != nil {
		return err
	}

	return ctx.JSONBytes(http.StatusOK, res)
}

func (c *controller) importer(ctx Context) error {
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
