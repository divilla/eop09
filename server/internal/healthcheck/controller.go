package healthcheck

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type controller struct {
	logger echo.Logger
}

func Controller(e *echo.Echo) {
	ctrl := &controller{
		logger: e.Logger,
		//ser: &service{},
	}

	e.GET("/healthcheck", ctrl.healthcheck)
}

func (c *controller) healthcheck(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
