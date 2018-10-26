package customerserv

import (
	"github.com/mperezguendulain/el-gallero-api-rest/daos"
	"github.com/mperezguendulain/el-gallero-api-rest/db"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2/bson"
)

var customerDAO *daos.CustomerDAO

// Inicializa el módulo
func init() {

	const customerCollection string = "customers"
	customerDAO = &daos.CustomerDAO{Collection: db.MgoSession.DB(db.DbName).C(customerCollection)}

}

// FindAllCustomers get all customers in Customers Collection
func FindAllCustomers() ([]models.Customer, error) {

	return customerDAO.FindAll()

}

// FindCustomerByID search for a customer by his id
func FindCustomerByID(uid string) (models.Customer, error) {

	return customerDAO.FindByID(uid)

}

// CreateCustomer create a record of a customer in db
func CreateCustomer(customer models.Customer) error {

	return customerDAO.Insert(customer)

}

// RemoveCustomerByID removes a customer in Customers Collection
func RemoveCustomerByID(customerID string) error {

	return customerDAO.RemoveByID(bson.ObjectIdHex(customerID))

}

// UpdateCustomer updates a customer in Customers Collection
func UpdateCustomer(uid string, customer models.Customer) error {

	customerID := bson.ObjectIdHex(uid)
	return customerDAO.Update(customerID, customer)

}
