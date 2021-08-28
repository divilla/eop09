package jsondecimals

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"regexp"
)

//Unquote removes quotes from quoted json decimal number values in order to present them back as JSON Numbers
//
//For setting path argument please refer to https://github.com/tidwall/gjson
func Unquote(json []byte, paths ...string) ([]byte, error) {
	var err error

	for _, path := range paths {
		result := gjson.GetBytes(json, path)
		switch result.Type {
		case gjson.String:
			json, err = unquoteValue(json, path, result)
			if err != nil {
				return nil, fmt.Errorf("error unquoting json decimal in path '%s' value '%s' with error: %w", path, result.Raw, err)
			}
		case gjson.JSON:
			if !result.IsArray() {
				continue
			}

			json, err = unquoteArray(json, path, result)
			if err != nil {
				return nil, fmt.Errorf("error unquoting json decimal in path '%s' array '%s' with error: %w", path, result.Raw, err)
			}
		default:
			continue
		}
	}

	return json, nil
}

func unquoteValue(json []byte, path string, result gjson.Result) ([]byte, error) {
	numberRegex := regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`)
	if result.Type != gjson.String || !numberRegex.MatchString(result.String()) {
		return json, nil
	}

	return sjson.SetRawBytes(json, path, []byte(result.String()))
}

func unquoteArray(json []byte, path string, result gjson.Result) ([]byte, error) {
	var err error

	json, err = sjson.SetRawBytes(json, path, []byte(`[]`))
	if err != nil {
		return nil, err
	}

	result.ForEach(func(key, value gjson.Result) bool {
		numberRegex := regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`)
		if value.Type != gjson.String || !numberRegex.MatchString(value.String()) {
			json, err = sjson.SetRawBytes(json, path + ".-1", []byte(value.Raw))
			return true
		}
		json, err = sjson.SetRawBytes(json, path + ".-1", []byte(value.String()))
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
