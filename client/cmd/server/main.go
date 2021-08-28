package main

import (
	"github.com/divilla/eop09/client/config"
	"github.com/divilla/eop09/client/internal/app"
	"github.com/divilla/eop09/client/internal/probe"
	"github.com/divilla/eop09/client/pkg/cecho"
	"github.com/divilla/eop09/client/pkg/cgrpc"
	"github.com/divilla/eop09/client/pkg/largejsonreader"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	_ = godotenv.Load(".env.devel")
}

func main() {
	// CPUProfile enables cpu profiling. Note: Default is CPU
	//defer profile.Start(profile.MemProfileHeap, profile.ProfilePath("/home/vito/go/projects/bootstrap/cmd/profile/")).Stop()

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	e.Use(cecho.CContext())
	e.HTTPErrorHandler = cecho.HTTPErrorHandler
	//e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	//	//Skipper:      middleware.DefaultSkipper,
	//	ErrorMessage: "request timeout, please try again",
	//	Timeout:      3*time.Second,
	//}))

	reader := largejsonreader.New(config.App.JsonDataFile)
	client := cgrpc.NewClient(config.App.GRPCServerAddress, e.Logger)
	defer client.Close()

	app.Controller(e, client, reader)
	probe.Controller(e)

	go func() {
		if err := e.Start(config.App.ServerAddress); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
