package app

import (
	"encoding/json"
	"github.com/divilla/eop09/entityproto"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	KeyPath = "key"
	CoordinatesPath = "coordinates"
)

func encodeEntityJson(key string, result *gjson.Result) (json.RawMessage, error) {
	if result.Type != gjson.JSON || !result.IsObject() {
		return nil, errors.New("invalid or malformed json file")
	}

	res, err := sjson.SetBytes([]byte(result.Raw), KeyPath, key)
	if err != nil {
		return nil, err
	}

	cordResult := result.Get(CoordinatesPath)
	if cordResult.Exists() && cordResult.IsArray() {
		res, err = sjson.SetRawBytes(res, CoordinatesPath, []byte(`[]`))
		if err != nil {
			return nil, err
		}
		cordResult.ForEach(func(key, value gjson.Result) bool {
			res, err = sjson.SetBytes(res, CoordinatesPath + ".-1", value.Raw)
			if err != nil {
				return false
			}
			return true
		})
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}

func decodeEntityJson(value json.RawMessage) (string, json.RawMessage, error) {
	key := gjson.GetBytes(value, KeyPath).String()
	res, err := sjson.DeleteBytes(value, KeyPath)
	if err != nil {
		return "", nil, err
	}

	result := gjson.GetBytes(value, CoordinatesPath)
	if result.Exists() && result.IsArray() {
		res, err = sjson.SetRawBytes(res, CoordinatesPath, []byte(`[]`))
		if err != nil {
			return "", nil, err
		}
		result.ForEach(func(key, value gjson.Result) bool {
			res, err = sjson.SetRawBytes(res, CoordinatesPath + ".-1", []byte(value.String()))
			if err != nil {
				return false
			}
			return true
		})
	}

	if err != nil {
		return "", nil, err
	}
	return key, res, nil
}

func parseImportResponse(is *entityproto.ImportResponse) (json.RawMessage, bool, error){
	res, err := sjson.SetBytes([]byte(`{}`), "success", is.GetSuccess())
	if err != nil {
		return nil, false, err
	}

	res, err = sjson.SetBytes(res, "rowsImported", is.GetRowsAffected())
	if err != nil {
		return nil, false, err
	}

	var k string
	var v json.RawMessage
	errRes := []byte(`{}`)
	if len(is.GetErrors()) > 4 {
		gjson.ParseBytes(is.GetErrors()).ForEach(func(key, value gjson.Result) bool {
			k, v, err = decodeEntityJson([]byte(value.Raw))
			if err != nil {
				return false
			}
			errRes, err = sjson.SetRawBytes(errRes, k, v)
			if err != nil {
				return false
			}
			return true
		})
		if err != nil {
			return nil, false, err
		}
		res, err = sjson.SetRawBytes(res, "errors", errRes)
		if err != nil {
			return nil, false, err
		}
	}

	return res, is.GetSuccess(), nil
}
