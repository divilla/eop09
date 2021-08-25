package main

import (
	"github.com/divilla/eop09/server/config"
	"github.com/divilla/eop09/server/internal/probe"
	"github.com/divilla/eop09/server/pkg/cmiddleware"
	"github.com/divilla/eop09/server/pkg/cmongo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"time"
	pb "github.com/divilla/eop09/crudproto"
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

	e.Logger.Info(os.Getenv("DSN"))
	mongo := cmongo.Init(os.Getenv("DSN"), e.Logger)
	col := mongo.Db().Collection("test")
	res, err := col.InsertOne(context.Background(), bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		panic(err)
	}
	e.Logger.Info(res.InsertedID)

	e.Use(cmiddleware.NewContext())
	e.HTTPErrorHandler = cmiddleware.HTTPErrorHandler
	//e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	//	//Skipper:      middleware.DefaultSkipper,
	//	ErrorMessage: "request timeout, please try again",
	//	Timeout:      3*time.Second,
	//}))

	probe.Controller(e)

	go func() {
		lis, err := net.Listen("tcp", config.App.ServerAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, &server{})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		//if err := e.Start(config.App.ServerAddress); err != nil && err != http.ErrServerClosed {
		//	e.Logger.Fatal("shutting down the server")
		//}
	}()

	//lis, err := net.Listen("tcp", serverPort)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//s := grpc.NewServer()
	//pb.RegisterPersistenceServer(s, &grpci.Server{})
	//if err = s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}


	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
