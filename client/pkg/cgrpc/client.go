package cgrpc

import (
	i "github.com/divilla/eop09/client/internal/interfaces"
	pb "github.com/divilla/eop09/entityproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
	"time"
)

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

func (c *Client) IsConnected() bool {
	if c.conn == nil {
		return false
	}

	state := c.conn.GetState()
	return state == connectivity.Ready || state == connectivity.Idle
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
	cp := grpc.ConnectParams{
		Backoff:           backoff.DefaultConfig,
		MinConnectTimeout: 3 * time.Second,
	}
	conn, err := grpc.Dial(c.serverAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithConnectParams(cp))
	if err != nil {
		c.logger.Panicf("gRPC client failed to connect: %w", err)
	} else {
		c.conn = conn
		c.RPCClient = pb.NewRPCClient(conn)
	}
}
