package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
)

type Context interface {
	echo.Context
	RequestContext() context.Context
	Body() string
	BodyBytes() []byte
	BodyJson() (gjson.Result, error)
	BodyMap() (map[string]interface{}, error)
	JSONBytes(code int, json []byte) error
	JSONString(code int, json string) error
	ParamInt64(name string, defaultValue int64) int64
	QueryParamInt64(name string, defaultValue int64) int64
}
