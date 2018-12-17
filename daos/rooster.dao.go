package daos

import (
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RoosterDAO is the DAO of Rooster
type RoosterDAO struct {
	Collection *mgo.Collection
}

// FindAll get all roosters in Roosters Collection
func (ud *RoosterDAO) FindAll() ([]models.Rooster, error) {

	roosters := make([]models.Rooster, 0)
	err := ud.Collection.Find(bson.M{}).All(&roosters)
	return roosters, err

}

// FindByID search for a rooster by his id
func (ud *RoosterDAO) FindByID(uid string) (models.Rooster, error) {

	var rooster models.Rooster
	err := ud.Collection.FindId(bson.ObjectIdHex(uid)).One(&rooster)
	return rooster, err

}

// Insert inserts a document in Roosters Collection
func (ud *RoosterDAO) Insert(rooster models.Rooster) error {

	err := ud.Collection.Insert(&rooster)
	return err

}

// RemoveByID removes a rooster in Roosters Collection
func (ud *RoosterDAO) RemoveByID(roosterID bson.ObjectId) error {

	err := ud.Collection.RemoveId(roosterID)
	return err

}

// Update updates a rooster in Roosters Collection
func (ud *RoosterDAO) Update(roosterID bson.ObjectId, rooster models.Rooster) error {

	err := ud.Collection.UpdateId(roosterID, &rooster)
	return err

}
