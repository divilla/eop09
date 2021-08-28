package jsondecimals

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//Quote quotes json decimal number values in order to preserve exact value after Unmarshalling
//Unmarshalling json decimal number values to float presents bad coding practice, cause float is not able to
//present all decimal numbers, resulting in invalid data stored in databases and calculated with.
//
//For setting path argument please refer to https://github.com/tidwall/gjson
func Quote(json []byte, paths ...string) ([]byte, error) {
	var err error

	for _, path := range paths {
		result := gjson.GetBytes(json, path)
		if !result.Exists() {
			continue
		}

		switch result.Type {
		case gjson.Number:
			json, err = quoteValue(json, path, result)
			if err != nil {
				return nil, fmt.Errorf("error quoting json decimal in path '%s' value '%s' with error: %w", path, result.Raw, err)
			}
		case gjson.JSON:
			if !result.IsArray() {
				continue
			}

			json, err = quoteArray(json, path, result)
			if err != nil {
				return nil, fmt.Errorf("error quoting json decimal in path '%s' array '%s' with error: %w", path, result.Raw, err)
			}
		default:
			continue
		}
	}

	return json, nil
}

func quoteValue(json []byte, path string, result gjson.Result) ([]byte, error) {
	return sjson.SetBytes(json, path, result.Raw)
}

func quoteArray(json []byte, path string, result gjson.Result) ([]byte, error) {
	var err error

	json, err = sjson.SetRawBytes(json, path, []byte(`[]`))
	if err != nil {
		return nil, err
	}

	result.ForEach(func(key, value gjson.Result) bool {
		if value.Type == gjson.Number {
			json, err = sjson.SetBytes(json, path+".-1", value.Raw)
		} else {
			json, err = sjson.SetRawBytes(json, path+".-1", []byte(value.Raw))
		}
		if err != nil {
			return false
		}

		return true
	})

	if err != nil {
		return nil, err
	}
	return json, nil
}
