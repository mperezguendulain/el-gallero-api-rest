package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Customer store the customer information
type Customer struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FullName   string        `json:"fullName" bson:"fullName"`
	Phone      string        `json:"phone" bson:"phone"`
	FiscalData FiscalData    `json:"fiscalData" bson:"fiscalData"`
}

// FiscalData represents the fiscal data of a customer.
type FiscalData struct {
	RFC          string `json:"rfc" bson:"rfc"`
	BusinessName string `json:"businessName" bson:"businessName"`
	Address      string `json:"address" bson:"address"`
	ZipCode      string `json:"zipCode" bson:"zipCode"`
	ContactPhone string `json:"contactPhone" bson:"contactPhone"`
}
