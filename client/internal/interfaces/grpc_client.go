package interfaces

import "github.com/divilla/eop09/entityproto"

type GRPCClient interface {
	entityproto.RPCClient
	IsConnected() bool
	State() string
}
type ImportClient entityproto.RPC_ImportClient
