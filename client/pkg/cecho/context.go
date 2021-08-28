package cecho

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strconv"
)

type (
	Context interface {
		echo.Context
		RequestContext() context.Context
		Body() string
		BodyBytes() []byte
		BodyJson() (gjson.Result, error)
		BodyMap() (map[string]interface{}, error)
		JSONBytes(code int, json []byte) error
		JSONString(code int, json string) error
		ParamInt64(name string, defaultValue int64) int64
	}

	ccontext struct {
		echo.Context
		//identity     *entity.Identity
	}
)

func CContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &ccontext{
				Context: c,
			}

			return next(ctx)
		}
	}
}

func (ctx *ccontext) RequestContext() context.Context {
	return ctx.Request().Context()
}

func (ctx *ccontext) Body() string {
	return string(ctx.BodyBytes())
}

func (ctx *ccontext) BodyBytes() []byte {
	body := ctx.Request().Body
	if body == nil {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}

	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes
}

func (ctx *ccontext) BodyJson() (gjson.Result, error) {
	res := gjson.Parse(ctx.Body())
	if res.Type != gjson.JSON {
		return res, echo.NewHTTPError(http.StatusBadRequest, "Invalid or malformed json request")
	}

	return res, nil
}

func (ctx *ccontext) BodyMap() (map[string]interface{}, error) {
	res, err := ctx.BodyJson()
	if err != nil {
		return nil, err
	}

	m, ok := res.Value().(map[string]interface{})
	if !ok {
		return m, echo.NewHTTPError(http.StatusBadRequest, "Unable to convert json to map")
	}

	return m, nil
}

func (ctx *ccontext) JSONBytes(code int, json []byte) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	_, err := ctx.Response().Write(json)
	if err != nil {
		return err
	}

	ctx.Response().Status = code
	return nil
}

func (ctx *ccontext) JSONString(code int, json string) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return ctx.String(code, json)
}

func (ctx *ccontext) ParamInt64(name string, defaultValue int64) int64 {
	value, err := strconv.ParseInt(ctx.Param(name), 10, 64)
	if err != nil || value < 1 {
		value = defaultValue
	}

	return value
}
