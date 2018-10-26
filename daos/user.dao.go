package daos

import (
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserDAO is the DAO of User
type UserDAO struct {
	Collection *mgo.Collection
}

// FindAll get all users in Users Collection
func (ud *UserDAO) FindAll() ([]models.User, error) {

	var users []models.User
	err := ud.Collection.Find(bson.M{}).Select(bson.M{"password": 0}).All(&users)
	return users, err

}

// FindByID search for a user by his id
func (ud *UserDAO) FindByID(uid string) (models.User, error) {

	var user models.User
	err := ud.Collection.FindId(bson.ObjectIdHex(uid)).Select(bson.M{"password": 0}).One(&user)
	return user, err

}

// Insert inserts a document in Users Collection
func (ud *UserDAO) Insert(user models.User) error {

	err := ud.Collection.Insert(&user)
	return err

}

// RemoveByID removes a user in Users Collection
func (ud *UserDAO) RemoveByID(userID bson.ObjectId) error {

	err := ud.Collection.RemoveId(userID)
	return err

}

// Update updates a user in Users Collection
func (ud *UserDAO) Update(userID bson.ObjectId, user models.User) error {

	err := ud.Collection.UpdateId(userID, &user)
	return err

}
