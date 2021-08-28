package app

import (
	i "github.com/divilla/eop09/client/internal/interfaces"
	. "github.com/divilla/eop09/client/pkg/cecho"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	g.GET("", H(ctrl.index))
	g.GET("/:key", H(ctrl.get))
	g.POST("", H(ctrl.create))
	g.PATCH("/:key", H(ctrl.patch))
	g.PUT("/:key", H(ctrl.put))
	g.DELETE("/:key", H(ctrl.delete))
	e.GET("/import", H(ctrl.importer))
}

func (c *controller) index(ctx i.Context) error {
	response, res, err := c.service.index(ctx.RequestContext(),
		ctx.QueryParamInt64("page", 1),
		ctx.QueryParamInt64("results", 30))
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("X-Pagination-Total-Count", strconv.FormatInt(res.TotalResults, 10))
	ctx.Response().Header().Set("X-Pagination-Page-Count", strconv.FormatInt(res.TotalPages, 10))
	ctx.Response().Header().Set("X-Pagination-Current-Page", strconv.FormatInt(res.Page, 10))
	ctx.Response().Header().Set("X-Pagination-Per-Page", strconv.FormatInt(res.PageSize, 10))
	return ctx.JSONBytes(http.StatusOK, response)
}

func (c *controller) get(ctx i.Context) error {
	res, err := c.service.get(ctx.RequestContext(), ctx.Param("key"))
	if err != nil {
		return err
	}

	return ctx.JSONBytes(http.StatusOK, res)
}

func (c *controller) create(ctx i.Context) error {
	req, err := ctx.BodyJson()
	if err != nil {
		return err
	}

	res, err := c.service.create(ctx.RequestContext(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *controller) patch(ctx i.Context) error {
	req, err := ctx.BodyJson()
	if err != nil {
		return err
	}

	res, err := c.service.patch(ctx.RequestContext(), ctx.Param("key"), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *controller) put(ctx i.Context) error {
	req, err := ctx.BodyJson()
	if err != nil {
		return err
	}

	res, err := c.service.put(ctx.RequestContext(), ctx.Param("key"), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

func (c *controller) delete(ctx i.Context) error {
	res, err := c.service.delete(ctx.RequestContext(), ctx.Param("key"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
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
