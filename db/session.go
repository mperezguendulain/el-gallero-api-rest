package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// MgoSession stores the main session of mongodb
var MgoSession *mgo.Session

// DbName stores tha database name
const DbName = "elGallero"

const urlDb = "localhost"

func init() {

	session, err := mgo.Dial(urlDb)
	if err != nil {
		log.Fatalln("Error >> Could not connect to the server:", err)
	}
	MgoSession = session

}
