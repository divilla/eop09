package main

import (
	"github.com/divilla/eop09/client/config"
	importer "github.com/divilla/eop09/client/internal/import"
	"github.com/divilla/eop09/client/internal/probe"
	"github.com/divilla/eop09/client/pkg/cmiddleware"
	"github.com/divilla/eop09/client/pkg/pgpool"
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

	pool := pgpool.Init(os.Getenv("DSN"), e.Logger)
	defer pool.Close()

	e.Use(cmiddleware.NewContext())
	e.HTTPErrorHandler = cmiddleware.HTTPErrorHandler
	//e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	//	//Skipper:      middleware.DefaultSkipper,
	//	ErrorMessage: "request timeout, please try again",
	//	Timeout:      3*time.Second,
	//}))

	importer.Controller(e)
	probe.Controller(e)

	go func() {
		if err := e.Start(config.App.ServerAddress); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
