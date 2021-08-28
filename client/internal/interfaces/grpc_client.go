package interfaces

import "github.com/divilla/eop09/crudproto"

type GRPCClient interface {
	entityproto.RPCClient
	IsConnected() bool
}
type ImportClient entityproto.RPC_ImportClient
