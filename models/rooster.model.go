package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Rooster store the rooster information
type Rooster struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Kind            string        `json:"kind" bson:"kind"`
	Age             int32         `json:"age" bson:"age"`
	Bloodline       string        `json:"bloodline" bson:"bloodline"`
	Characteristics string        `json:"characteristics" bson:"characteristics"`
	Colors          RoosterColors `json:"colors" bson:"colors"`
}

// RoosterColors store the colors information of a rooster.
type RoosterColors struct {
	Bolilla       string `json:"bolilla" bson:"bolilla"`
	LegColour     string `json:"legColour" bson:"legColour"`
	ColorOfGuides string `json:"colorOfGuides" bson:"colorOfGuides"`
}
