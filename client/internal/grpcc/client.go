package grpcc

import (
	"fmt"
	"github.com/divilla/eop09/client/internal/interfaces"
	"github.com/divilla/eop09/crudproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	serverAddress string
	conn          *grpc.ClientConn
	client        crudproto.RPCClient
	logger        interfaces.Logger
}

func NewClient(serverAddress string, logger interfaces.Logger) (*Client, error) {
	c := &Client{
		serverAddress: serverAddress,
		logger:        logger,
	}
	err := c.dial()

	return c, err
}

func (c *Client) List(ctx context.Context, listRequest *crudproto.ListRequest) (*crudproto.ListResponse, error) {
	return c.client.List(ctx, listRequest)
}

func (c *Client) Import(ctx context.Context) (crudproto.RPC_ImportClient, error) {
	return c.client.Import(ctx)
}

func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("gRPC client failed to close: %w", err)
	}

	return nil
}

func (c *Client) dial() error {
	conn, err := grpc.Dial(c.serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		c.logger.Fatalf("gRPC client failed to connect: %w", err)
	} else {
		c.conn = conn
		c.client = crudproto.NewRPCClient(conn)
	}

	return err
}
