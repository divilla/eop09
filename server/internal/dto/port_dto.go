package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
	"strconv"
)

type (
	PortDto struct {
		Id          string                 `json:"key" bson:"_id"`
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

	Errors []byte
)

func (p *PortDto) Validate() validation.Errors {
	err := validation.ValidateStruct(p,
		validation.Field(&p.Id, validation.Required.Error("required"),
			validation.Match(regexp.MustCompile("^[A-Z]{5}$")).Error("must be 5 uppercase letters word")),
		validation.Field(&p.Name, validation.Required.Error("required")),
		validation.Field(&p.City, validation.Required.Error("required")),
		validation.Field(&p.Country, validation.Required.Error("required")),
		validation.Field(&p.Coordinates, validation.By(func(value interface{}) error {
			if len(p.Coordinates) != 2 {
				return errors.New("two decimal numbers required")
			}
			es := validation.Errors{}
			for k, v := range p.Coordinates {
				reg := regexp.MustCompile("^[+-]?(\\d*\\.)?\\d+$")
				if !reg.MatchString(v.String()) {
					es[strconv.Itoa(k)] = errors.New("not a decimal number")
				}
			}
			if len(es) > 0 {
				return es
			}
			return nil
		})),
		validation.Field(&p.Timezone, validation.Required.Error("required")),
		validation.Field(&p.Code, validation.Required.Error("required"),
			validation.Match(regexp.MustCompile("^\\d{5}$")).Error("must be 5 digits number")),
	)

	return err.(validation.Errors)
}

func (e Errors) Error() string {
	return string(e)
}
