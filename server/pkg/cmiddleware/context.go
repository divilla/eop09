package cmiddleware

import (
	"bytes"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
)

type Context struct {
	echo.Context
	conn *pgxpool.Conn
	//identity     *entity.Identity
}

func NewContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &Context{
				Context: c,
			}

			return next(ctx)
		}
	}
}

func (ctx *Context) RequestContext() context.Context {
	return ctx.Request().Context()
}

//func (ctx *Context) Identity() (*entity.Identity, error) {
//	if ctx.identity == nil {
//		return nil, cerrors.NotLoggedInErr.Wrap(cerrors.NewCode("pkg.cecho.context_middleware.Context.Identity()", ""))
//	}
//
//	return ctx.identity, nil
//}

func (ctx *Context) Body() string {
	var bodyBytes []byte
	var err error

	if ctx.Request().Body != nil {
		bodyBytes, err = ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			panic(err)
		}
	}

	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return string(bodyBytes)
}

func (ctx *Context) BodyJson() ([]byte, error) {
	res := gjson.Parse(ctx.Body())
	if res.Type != gjson.JSON {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid or malformed json request")
	}

	return []byte(res.Raw), nil
}

func (ctx *Context) BodyGJson() (*gjson.Result, error) {
	res := gjson.Parse(ctx.Body())
	if res.Type != gjson.JSON {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid or malformed json request")
	}

	return &res, nil
}

func (ctx *Context) BodyMap() (map[string]interface{}, error) {
	res, err := ctx.BodyGJson()
	if err != nil {
		return nil, err
	}

	m, ok := res.Value().(map[string]interface{})
	if !ok {
		return m, echo.NewHTTPError(http.StatusBadRequest, "Unable to convert json to map")
	}

	return m, nil
}

func (ctx *Context) JSONS(code int, json string) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return ctx.String(code, json)
}
