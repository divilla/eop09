package app

import (
	"encoding/json"
	"github.com/divilla/eop09/entityproto"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"net/http"
)

const (
	keyPath         = "key"
	coordinatesPath = "coordinates"
)

//incoming json object must be parsed in order to make it ready for unmarshalling
//object key is sent in a form of property name, so it needs to be added to value object
//coordinates are sent in a form of number, which, if imported as float might mutate it's value
//therefore coordinates must be converted to string and then unmarshalled to decimal to preserve their original value
func encodeEntityJson(result *gjson.Result) (json.RawMessage, error) {
	var key string
	var value *gjson.Result
	var bValue json.RawMessage

	result.ForEach(func(k, v gjson.Result) bool {
		key = k.String()
		value = &v
		return false
	})

	if key == "" || value == nil || !value.Exists() || !value.IsObject() {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid or malformed json request")
	}

	bValue = []byte(value.Raw)
	return encodeEntityKeyValue(key, &bValue)
}

func encodeEntityKeyValue(key string, value *json.RawMessage) (json.RawMessage, error) {
	// as key comes in form of property name, the following line will add it to value object
	res, err := sjson.SetBytes(*value, keyPath, key)
	if err != nil {
		panic(err)
	}

	// in order to preserve original decimal value, coordinates are quoted and turned into string
	// later they can be converted to and handled as primitive.Decimal128
	cordResult := gjson.GetBytes(*value, coordinatesPath)
	if cordResult.Exists() && cordResult.IsArray() {
		res, err = sjson.SetRawBytes(res, coordinatesPath, []byte(`[]`))
		if err != nil {
			panic(err)
		}
		cordResult.ForEach(func(key, value gjson.Result) bool {
			res, err = sjson.SetBytes(res, coordinatesPath+".-1", value.Raw)
			if err != nil {
				panic(err)
			}
			return true
		})
	}

	return res, nil
}

func decodeEntityJson(value json.RawMessage) (string, json.RawMessage, error) {
	key := gjson.GetBytes(value, keyPath).String()
	res, err := sjson.DeleteBytes(value, keyPath)
	if err != nil {
		return "", nil, err
	}

	result := gjson.GetBytes(value, coordinatesPath)
	if result.Exists() && result.IsArray() {
		res, err = sjson.SetRawBytes(res, coordinatesPath, []byte(`[]`))
		if err != nil {
			return "", nil, err
		}
		result.ForEach(func(key, value gjson.Result) bool {
			res, err = sjson.SetRawBytes(res, coordinatesPath+".-1", []byte(value.String()))
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

func parseImportResponse(is *entityproto.ImportResponse) (json.RawMessage, bool, error) {
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
