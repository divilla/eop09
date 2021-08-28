package interfaces

import "encoding/json"

type JsonReader interface {
	Read(index *uint64, key *string, value *json.RawMessage) error
	Reset()
	Close() error
}
