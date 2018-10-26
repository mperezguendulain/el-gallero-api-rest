package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User store the information of a User
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password,omitempty" bson:"password"`
	FullName string        `json:"fullName" bson:"fullName"`
	Type     string        `json:"type" bson:"type"`
	Address  string        `json:"address" bson:"address"`
	Contact  UserContact   `json:"contact" bson:"contact"`
}

// UserContact represents the contact data of a User.
type UserContact struct {
	Email     string `json:"email" bson:"email"`
	Cellphone string `json:"cellphone" bson:"cellphone"`
	Landline  string `json:"landLine" bson:"landLine"`
}
