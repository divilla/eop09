package cmongo

import (
	"encoding/json"
	"github.com/tidwall/sjson"
)

type JsonError json.RawMessage

func NewJsonError(code int, message string) JsonError {
	j, _ := sjson.SetBytes([]byte(`{}`), "code", code)
	j, _ = sjson.SetBytes(j, "message", message)

	return j
}

func (j JsonError) Error() string {
	return string(j)
}
