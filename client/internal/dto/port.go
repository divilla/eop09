package dto

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type (
	PortDto struct {
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

func (e *PortDto) Marshall() ([]byte, error) {
	return json.Marshal(e)
}

func (e *PortDto) Unmarshal(b []byte) error {
	return json.Unmarshal(b, e)
}
