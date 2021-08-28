package cecho

import (
	i "github.com/divilla/eop09/client/internal/interfaces"
	"github.com/labstack/echo/v4"
)

type HandlerFunc func(ctx i.Context) error

func H(hf HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*ccontext)
		return hf(ctx)
	}
}
