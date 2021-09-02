package cgrpc

import (
	i "github.com/divilla/eop09/client/internal/interfaces"
	pb "github.com/divilla/eop09/entityproto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
	_ "google.golang.org/grpc/health"
	"google.golang.org/grpc/keepalive"
	"time"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

var cp = grpc.ConnectParams{
	Backoff:           backoff.DefaultConfig,
	MinConnectTimeout: 3 * time.Second,
}

var sc = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
}`

type Client struct {
	pb.RPCClient
	serverAddress string
	conn          *grpc.ClientConn
	logger        i.Logger
}

func NewClient(serverAddress string, logger i.Logger) *Client {
	c := &Client{
		serverAddress: serverAddress,
		logger:        logger,
	}
	go c.dial()

	return c
}

func (c *Client) Ping() error {
	if c.conn == nil {
		return errors.Errorf("not connected to gRPC server: %s", c.serverAddress)
	}

	state := c.conn.GetState()
	if state != connectivity.Ready && state != connectivity.Idle {
		return errors.Errorf("lost connection to gRPC server: %s", c.serverAddress)
	}

	return nil
}

func (c *Client) Close() {
	if c.conn == nil || c.conn.GetState() == connectivity.Shutdown {
		return
	}

	err := c.conn.Close()
	if err != nil {
		c.logger.Panicf("gRPC client failed to close: %w", err)
	}
}

func (c *Client) dial() {
	conn, err := grpc.Dial(c.serverAddress,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithBlock(),
		grpc.WithConnectParams(cp),
		grpc.WithDefaultServiceConfig(sc))
	if err != nil {
		c.logger.Panicf("gRPC client failed to connect to '%s' with error: %w", c.serverAddress, err)
	} else {
		c.conn = conn
		c.RPCClient = pb.NewRPCClient(conn)
	}
}
