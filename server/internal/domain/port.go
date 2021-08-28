package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Port struct {
		Id          primitive.ObjectID     `json:"-" bson:"_id"`
		Key         string                 `json:"-" bson:"key"`
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
