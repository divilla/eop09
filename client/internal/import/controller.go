package importer

import (
	"github.com/divilla/eop09/client/config"
	"github.com/divilla/eop09/client/internal/domain"
	interfaces2 "github.com/divilla/eop09/client/internal/interfaces"
	jsonfilereader "github.com/divilla/eop09/client/pkg/jReader"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type controller struct {
	logger interfaces2.Logger
}

func Controller(e *echo.Echo) {
	ctrl := &controller{
		logger: e.Logger,
		//ser: &service{},
	}

	e.GET("/import", ctrl.importer)
}

func (c *controller) importer(ctx echo.Context) error {
	var port domain.Port

	jfr := jsonfilereader.Init(config.App.JsonDataFile, c.logger)
	jfr.Parse(&port, func(wg *sync.WaitGroup, parser interface{}, err error) {
		if err != nil {
			c.logger.Errorf("Unable to parse json: ", err)
		}
		c.logger.Info(parser)

		wg.Done()
	})
	jfr.Close()

	return ctx.NoContent(http.StatusOK)
}
