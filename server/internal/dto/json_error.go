package dto

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tidwall/sjson"
)

type JsonErrors json.RawMessage

func NewValidationErrors(ve []byte) *JsonErrors {
	j := new(JsonErrors)
	*j, _ = sjson.SetBytes([]byte(`{}`), "code", 422)
	*j, _ = sjson.SetRawBytes(*j, "errors", ve)

	return j
}

func NewJsonError(code int, message ...string) *JsonErrors {
	j := new(JsonErrors)
	*j, _ = sjson.SetBytes([]byte(`{}`), "code", code)
	if len(message) > 0 {
		*j, _ = sjson.SetBytes(*j, "message", message[0])
	}

	return j
}

func (j *JsonErrors) Message(message string) *JsonErrors {
	*j, _ = sjson.SetBytes(*j, "message", message)
	return j
}

func (j *JsonErrors) Errors(e validation.Errors) *JsonErrors {
	eb, _ := e.MarshalJSON()
	*j, _ = sjson.SetRawBytes(*j, "errors", eb)
	return j
}

func (j *JsonErrors) Error() string {
	return string(*j)
}
