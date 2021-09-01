package app

import (
	i "github.com/divilla/eop09/client/internal/interfaces"
	ce "github.com/divilla/eop09/client/pkg/cecho"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type controller struct {
	service *service
	logger  i.Logger
}

//Controller builds app main controller
func Controller(e *echo.Echo, client i.GRPCClient, reader i.JsonReader) {
	ctrl := &controller{
		service: newService(client, reader, e.Logger),
		logger:  e.Logger,
	}

	// RFC REST setup
	g := e.Group("/ports")
	g.GET("", ce.H(ctrl.index))
	g.GET("/:key", ce.H(ctrl.get))
	g.POST("", ce.H(ctrl.create))
	g.PATCH("/:key", ce.H(ctrl.patch))
	g.PUT("/:key", ce.H(ctrl.put))
	g.DELETE("/:key", ce.H(ctrl.delete))

	// import is not part of ports group, because of conflict with possible 'import' key
	e.GET("/import", ce.H(ctrl.importer))
}

//index is used to list Ports json.
//It accepts query parameters 'page' - page number of result set, 'results' - number of results displayed per page
func (c *controller) index(ctx i.Context) error {
	response, res, err := c.service.index(ctx.RequestContext(),
		ctx.QueryParamInt64("page", 1),
		ctx.QueryParamInt64("results", 30))
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("X-Pagination-Total-Count", strconv.FormatInt(res.TotalCount, 10))
	ctx.Response().Header().Set("X-Pagination-Page-Count", strconv.FormatInt(res.PageCount, 10))
	ctx.Response().Header().Set("X-Pagination-Current-Page", strconv.FormatInt(res.CurrentPage, 10))
	ctx.Response().Header().Set("X-Pagination-Per-Page", strconv.FormatInt(res.PerPage, 10))

	return ctx.JSONBytes(http.StatusOK, response)
}

//get is used to fetch single Port.
//It accepts Port 'key' parameter in the end of url
func (c *controller) get(ctx i.Context) error {
	res, err := c.service.get(ctx.RequestContext(), ctx.Param("key"))
	if err != nil {
		return err
	}

	return ctx.JSONBytes(http.StatusOK, res)
}

//create is used to register new Port
//It accepts JSON object in form of {"key": {"name": "Some Name", ...}}.
//See /data/ports.json
func (c *controller) create(ctx i.Context) error {
	req, err := ctx.BodyGJson()
	if err != nil {
		return err
	}

	res, err := c.service.create(ctx.RequestContext(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

//patch is used to modify some properties of existing Port
//It accepts Port 'key' parameter in the end of url and JSON object in form of {"key": {"name": "Some Name", ...}}.
//See /data/ports.json
func (c *controller) patch(ctx i.Context) error {
	req, err := ctx.BodyGJson()
	if err != nil {
		return err
	}

	res, err := c.service.patch(ctx.RequestContext(), ctx.Param("key"), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

//put is used to replace existing Port with new set of values
//It accepts Port 'key' parameter in the end of url and JSON object in form of {"key": {"name": "Some Name", ...}}.
//See /data/ports.json
func (c *controller) put(ctx i.Context) error {
	req, err := ctx.BodyGJson()
	if err != nil {
		return err
	}

	res, err := c.service.put(ctx.RequestContext(), ctx.Param("key"), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

//delete is used to remove existing Port from database
//It accepts Port 'key' parameter in the end of url
func (c *controller) delete(ctx i.Context) error {
	res, err := c.service.delete(ctx.RequestContext(), ctx.Param("key"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)
}

//importer is used to import values from file /data/ports.json
//It reads the file buffering key and value of each entry and sending them to gRPC upstream
//It upserts, meaning if it's new key it will insert new value, if it's existing key it will replace existing value
func (c *controller) importer(ctx i.Context) error {
	res, success, err := c.service.importer(ctx.RequestContext())
	if err != nil {
		return err
	}

	status := http.StatusOK
	if !success {
		status = http.StatusBadRequest
	}

	return ctx.JSON(status, res)
}
