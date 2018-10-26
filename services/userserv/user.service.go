package userserv

import (
	"github.com/mperezguendulain/el-gallero-api-rest/daos"
	"github.com/mperezguendulain/el-gallero-api-rest/db"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2/bson"
)

var userDao *daos.UserDAO

// Inicializa el módulo
func init() {

	const userCollection string = "users"
	userDao = &daos.UserDAO{Collection: db.MgoSession.DB(db.DbName).C(userCollection)}

}

// FindAllUsers get all users in Users Collection
func FindAllUsers() ([]models.User, error) {

	return userDao.FindAll()

}

// FindUserByID search for a user by his id
func FindUserByID(uid string) (models.User, error) {

	return userDao.FindByID(uid)

}

// CreateUser create a record of a user in db
func CreateUser(user models.User) error {

	return userDao.Insert(user)

}

// RemoveUserByID removes a user in Users Collection
func RemoveUserByID(userID string) error {

	return userDao.RemoveByID(bson.ObjectIdHex(userID))

}

// UpdateUser updates a user in Users Collection
func UpdateUser(uid string, user models.User) error {

	userID := bson.ObjectIdHex(uid)
	return userDao.Update(userID, user)

}
