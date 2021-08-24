package probe

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

	g := e.Group("/probe")
	g.GET("/liveness", ctrl.liveness)
}

func (c *controller) liveness(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
