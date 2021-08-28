package interfaces

import "encoding/json"

type JsonReader interface {
	Start() error
	Read(index *uint64, key *string, value *json.RawMessage) error
	Close() error
}
