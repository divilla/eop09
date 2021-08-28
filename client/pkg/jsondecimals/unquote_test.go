package jsondecimals

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
)

var jsonQuoted = `{
	"name": "Ajman",
	"city": "Ajman",
	"country": "United Arab Emirates",
	"alias": [],
	"regions": [],
	"nr": "15",
	"nrs": [
		"11.11", 
		"22.22", 
		null,
		"null"
	],
	"coordinates": [
	  55.5136433,
	  25.4052165
	],
	"province": "Ajman",
	"timezone": "Asia/Dubai",
	"unlocs": [
	  "AEAJM"
	],
	"code": "52000"
}`

func TestUnquote(t *testing.T) {
	json, err := Unquote([]byte(jsonQuoted), "nr", "nrs", "coordinates", "code")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, float64(15), gjson.GetBytes(json, "nr").Value(), "nr 15 should not be quoted")
	assert.Equal(t, 11.11, gjson.GetBytes(json, "nrs.0").Value(), "nrs.0 11.11 should not be quoted")
	assert.Equal(t, 22.22, gjson.GetBytes(json, "nrs.1").Value(), "nr 15 should not be quoted")
	assert.Equal(t, nil, gjson.GetBytes(json, "nrs.2").Value(), "nr null should not be quoted")
}
