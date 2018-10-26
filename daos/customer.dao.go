package daos

import (
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CustomerDAO is the DAO of Customer
type CustomerDAO struct {
	Collection *mgo.Collection
}

// FindAll get all customers in Customers Collection
func (cd *CustomerDAO) FindAll() ([]models.Customer, error) {

	var customers []models.Customer
	err := cd.Collection.Find(bson.M{}).All(&customers)
	return customers, err

}

// FindByID search for a customer by his id
func (cd *CustomerDAO) FindByID(customerID string) (models.Customer, error) {

	var customer models.Customer
	err := cd.Collection.FindId(bson.ObjectIdHex(customerID)).One(&customer)
	return customer, err

}

// Insert inserts a document in customers Collection
func (cd *CustomerDAO) Insert(customer models.Customer) error {

	err := cd.Collection.Insert(&customer)
	return err

}

// RemoveByID removes a customer in customers Collection
func (cd *CustomerDAO) RemoveByID(customerID bson.ObjectId) error {

	err := cd.Collection.RemoveId(customerID)
	return err

}

// Update updates a customer in customers Collection
func (cd *CustomerDAO) Update(customerID bson.ObjectId, customer models.Customer) error {

	err := cd.Collection.UpdateId(customerID, &customer)
	return err

}
