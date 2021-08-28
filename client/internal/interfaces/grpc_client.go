package interfaces

import "github.com/divilla/eop09/crudproto"

type GRPCClient interface {
	crudproto.RPCClient
	IsConnected() bool
}
type ImportClient crudproto.RPC_ImportClient
