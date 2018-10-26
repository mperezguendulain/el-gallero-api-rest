package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Provider store the provider information
type Provider struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FullName  string        `json:"fullName" bson:"fullName"`
	Products  []string      `json:"products" bson:"products"`
	Cellphone string        `json:"cellphone" bson:"cellphone"`
	Landline  string        `json:"landline" bson:"landline"`
	Address   string        `json:"address" bson:"address"`
}
