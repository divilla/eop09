package healthcheck

import (
	i "github.com/divilla/eop09/client/internal/interfaces"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type controller struct {
	client i.GRPCClient
	logger echo.Logger
}

func Controller(e *echo.Echo, client i.GRPCClient) {
	ctrl := &controller{
		client: client,
		logger: e.Logger,
	}

	e.GET("/healthcheck", ctrl.healthcheck)
}

func (c *controller) healthcheck(ctx echo.Context) error {
	var err error
	for i:=0; i<3; i++ {
		err = c.client.Ping()
		if  err == nil {
			return ctx.NoContent(http.StatusOK)
		}
		time.Sleep(time.Second)
	}

	c.logger.Fatal(err)
	return ctx.NoContent(http.StatusGone)
}
