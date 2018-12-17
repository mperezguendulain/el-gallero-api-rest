package roosterserv

import (
	"github.com/mperezguendulain/el-gallero-api-rest/daos"
	"github.com/mperezguendulain/el-gallero-api-rest/db"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2/bson"
)

var roosterDAO *daos.RoosterDAO

// Init the module
func init() {

	const roosterCollection string = "roosters"
	roosterDAO = &daos.RoosterDAO{Collection: db.MgoSession.DB(db.DbName).C(roosterCollection)}

}

// FindAllRoosters get all roosters in Roosters Collection
func FindAllRoosters() ([]models.Rooster, error) {

	return roosterDAO.FindAll()

}

// FindRoosterByID search for a rooster by his id
func FindRoosterByID(uid string) (models.Rooster, error) {

	return roosterDAO.FindByID(uid)

}

// CreateRooster create a record of a rooster in db
func CreateRooster(rooster models.Rooster) error {

	return roosterDAO.Insert(rooster)

}

// RemoveRoosterByID removes a rooster in Roosters Collection
func RemoveRoosterByID(roosterID string) error {

	return roosterDAO.RemoveByID(bson.ObjectIdHex(roosterID))

}

// UpdateRooster updates a rooster in Roosters Collection
func UpdateRooster(uid string, rooster models.Rooster) error {

	roosterID := bson.ObjectIdHex(uid)
	return roosterDAO.Update(roosterID, rooster)

}
