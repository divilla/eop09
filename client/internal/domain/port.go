package domain

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type (
	Ider interface {
		SetId(id string)
	}

	Port struct {
		Id          string            `json:"id"`
		Name        string            `json:"name"`
		City        string            `json:"city"`
		Country     string            `json:"country"`
		Alias       []string          `json:"alias"`
		Regions     []string          `json:"regions"`
		Coordinates []decimal.Decimal `json:"coordinates"`
		Province    string            `json:"province"`
		Timezone    string            `json:"timezone"`
		Unlocs      []string          `json:"unlocs"`
		Code        string            `json:"code"`
	}
)

func (e *Port) SetId(id string) {
	e.Id = id
}

func (e *Port) Marshall() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Port) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}
