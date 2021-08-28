package dto

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tidwall/sjson"
)

type JsonError json.RawMessage

func NewValidationErrors(ve []byte) *JsonError {
	j := new(JsonError)
	*j, _ = sjson.SetBytes([]byte(`{}`), "code", 422)
	*j, _ = sjson.SetRawBytes(*j, "errors", ve)

	return j
}

func NewJsonError(code int, message ...string) *JsonError {
	j := new(JsonError)
	*j, _ = sjson.SetBytes([]byte(`{}`), "code", code)
	if len(message) > 0 {
		*j, _ = sjson.SetBytes(*j, "message", message[0])
	}

	return j
}

func (j *JsonError) Message(message string) *JsonError {
	*j, _ = sjson.SetBytes(*j, "message", message)
	return j
}

func (j *JsonError) Errors(e validation.Errors) *JsonError {
	eb, _ := e.MarshalJSON()
	*j, _ = sjson.SetRawBytes(*j, "errors", eb)
	return j
}

func (j *JsonError) Error() string {
	return string(*j)
}
