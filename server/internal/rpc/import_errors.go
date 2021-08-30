package rpc

import (
	"encoding/json"
	"github.com/tidwall/sjson"
)

const KeyPath = "key"

type importErrors json.RawMessage

func newJsonErrors() *importErrors {
	var i importErrors
	i = []byte(`[]`)

	return &i
}

func (j *importErrors) Add(value json.RawMessage) error {
	var err error
	*j, err = sjson.SetRawBytes(*j, `-1`, value)

	return err
}

func (j *importErrors) Errors() json.RawMessage {
	if len(*j) == 0 {
		return nil
	}
	return json.RawMessage(*j)
}
