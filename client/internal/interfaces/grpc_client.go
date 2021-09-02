package interfaces

import "github.com/divilla/eop09/entityproto"

type GRPCClient interface {
	entityproto.RPCClient
	Ping() error
}
type ImportClient entityproto.RPC_ImportClient
