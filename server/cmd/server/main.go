package main

import (
	"flag"
	pb "github.com/divilla/eop09/entityproto"
	"github.com/divilla/eop09/server/internal/config"
	"github.com/divilla/eop09/server/internal/rpc"
	"github.com/divilla/eop09/server/pkg/cmongo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

var flagConfig = flag.String("mode", "local", "select config file")

func main() {
	// CPUProfile enables cpu profiling. Note: Default is CPU
	//defer profile.Start(profile.MemProfileHeap, profile.ProfilePath("/home/vito/go/projects/bootstrap/cmd/profile/")).Stop()

	flag.Parse()
	config.Init(*flagConfig)

	logger := log.New("eop09")
	logger.SetLevel(log.INFO)

	mongo := cmongo.Init(viper.GetString("ports_dsn"), logger)
	rep := cmongo.NewRepository(mongo, "port")

	lis, err := net.Listen("tcp", viper.GetString("server_address"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRPCServer(s, rpc.NewServer(rep, logger))
	logger.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
