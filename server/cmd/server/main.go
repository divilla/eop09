package main

import (
	"context"
	"flag"
	pb "github.com/divilla/eop09/entityproto"
	"github.com/divilla/eop09/server/internal/config"
	"github.com/divilla/eop09/server/internal/rpc"
	"github.com/divilla/eop09/server/pkg/cmongo"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

var flagConfig = flag.String("mode", "local", "select config file")

func main() {
	// CPUProfile enables cpu profiling. Note: Default is CPU
	//defer profile.Start(profile.MemProfileHeap, profile.ProfilePath("/home/vito/go/projects/bootstrap/cmd/profile/")).Stop()

	flag.Parse()
	config.Init(*flagConfig)

	zapLogger, _ := zap.NewProduction()
	defer func(zapLogger *zap.Logger) {
		err := zapLogger.Sync()
		if err != nil {
			panic(err)
		}
	}(zapLogger)

	logger := log.New("eop09")
	logger.SetLevel(log.INFO)

	mongo := cmongo.Init(viper.GetString("ports_dsn"), logger)
	rep := cmongo.NewRepository(mongo, "port")

	lis, err := net.Listen("tcp", ":" + viper.GetString("server_port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(kaep),
		grpc.KeepaliveParams(kasp),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zapLogger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
		)),
	)
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(s, healthcheck)
	pb.RegisterRPCServer(s, rpc.NewServer(rep, logger))

	go healthCheck(healthcheck, mongo)

	logger.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}

func healthCheck(healthcheck *health.Server, client *cmongo.CMongo) {
	ctx := context.TODO()
	healthcheck.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	for {
		time.Sleep(6*time.Second)

		err := client.Ping(ctx)
		if err != nil {
			healthcheck.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
		} else {
			healthcheck.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
		}
	}
}
