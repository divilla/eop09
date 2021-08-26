package main

import (
	"github.com/divilla/eop09/crudproto"
	"github.com/divilla/eop09/server/config"
	"github.com/divilla/eop09/server/internal/rpc"
	"github.com/divilla/eop09/server/pkg/cmongo"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
)

func init() {
	_ = godotenv.Load(".env.devel")
}

func main() {
	// CPUProfile enables cpu profiling. Note: Default is CPU
	//defer profile.Start(profile.MemProfileHeap, profile.ProfilePath("/home/vito/go/projects/bootstrap/cmd/profile/")).Stop()

	logger := log.New("server")
	logger.SetLevel(log.INFO)

	mongo := cmongo.Init(os.Getenv("DSN"), logger)
	col := mongo.Db().Collection("test")
	res, err := col.InsertOne(context.Background(), bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		panic(err)
	}
	logger.Info(res.InsertedID)

	lis, err := net.Listen("tcp", config.App.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	crudproto.RegisterRPCServer(s, rpc.NewServer(logger))
	logger.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
