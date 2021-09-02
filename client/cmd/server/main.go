package main

import (
	"flag"
	"github.com/divilla/eop09/client/internal/app"
	"github.com/divilla/eop09/client/internal/config"
	"github.com/divilla/eop09/client/internal/healthcheck"
	"github.com/divilla/eop09/client/pkg/cecho"
	"github.com/divilla/eop09/client/pkg/cgrpc"
	"github.com/divilla/eop09/client/pkg/largejsonreader"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var flagConfig = flag.String("mode", "local", "select config file")

func main() {
	// CPUProfile enables cpu profiling. Note: Default is CPU
	//defer profile.Start(profile.MemProfileHeap, profile.ProfilePath("/home/vito/go/projects/bootstrap/cmd/profile/")).Stop()

	flag.Parse()
	config.Init(*flagConfig)

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

	reader := largejsonreader.New(viper.GetString("json_data_file"))
	client := cgrpc.NewClient(viper.GetString("ports_grpc"), e.Logger)
	defer client.Close()

	app.Controller(e, client, reader)
	healthcheck.Controller(e, client)

	go func() {
		if err := e.Start(":" + viper.GetString("server_port")); err != nil && err != http.ErrServerClosed {
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
