package dto

import (
	"encoding/json"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

type (
	Port struct {
		Key         string                 `json:"key" bson:"key"`
		Name        string                 `json:"name" bson:"name"`
		City        string                 `json:"city" bson:"city"`
		Country     string                 `json:"country" bson:"country"`
		Alias       []string               `json:"alias" bson:"alias"`
		Regions     []string               `json:"regions" bson:"regions"`
		Coordinates []primitive.Decimal128 `json:"coordinates" bson:"coordinates"`
		Province    string                 `json:"province" bson:"province"`
		Timezone    string                 `json:"timezone" bson:"timezone"`
		Unlocs      []string               `json:"unlocs" bson:"unlocs"`
		Code        string                 `json:"code" bson:"code"`
	}
)

func (p *Port) Validate() val.Errors {
	err := val.ValidateStruct(p,
		val.Field(&p.Key, val.Required.Error("required"),
			val.Match(regexp.MustCompile("^[A-Z0-9]{5}$")).Error("must be 5 uppercase letters word")),
		val.Field(&p.Name, val.Required.Error("required")),
		val.Field(&p.City, val.Required.Error("required")),
		val.Field(&p.Country, val.Required.Error("required")),
		val.Field(&p.Coordinates, val.Length(2,2).Error("exactly two decimal values required")),
		val.Field(&p.Code, val.Match(regexp.MustCompile("^\\d{5}$")).Error("must be 5 digits number")),
	)

	if err == nil {
		return nil
	}
	return err.(val.Errors)
}

func (p *Port) ValidateAndMarshal() (json.RawMessage, error) {
	errs := p.Validate()
	if errs == nil {
		return nil, nil
	}

	return errs.MarshalJSON()
}
