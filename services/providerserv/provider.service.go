package providerserv

import (
	"github.com/mperezguendulain/el-gallero-api-rest/daos"
	"github.com/mperezguendulain/el-gallero-api-rest/db"
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2/bson"
)

var providerDAO *daos.ProviderDAO

// Init the module
func init() {

	const providerCollection string = "providers"
	providerDAO = &daos.ProviderDAO{Collection: db.MgoSession.DB(db.DbName).C(providerCollection)}

}

// FindAllProviders get all providers in Providers Collection
func FindAllProviders() ([]models.Provider, error) {

	return providerDAO.FindAll()

}

// FindProviderByID search for a provider by his id
func FindProviderByID(uid string) (models.Provider, error) {

	return providerDAO.FindByID(uid)

}

// CreateProvider create a record of a provider in db
func CreateProvider(provider models.Provider) error {

	return providerDAO.Insert(provider)

}

// RemoveProviderByID removes a provider in Providers Collection
func RemoveProviderByID(providerID string) error {

	return providerDAO.RemoveByID(bson.ObjectIdHex(providerID))

}

// UpdateProvider updates a provider in Providers Collection
func UpdateProvider(uid string, provider models.Provider) error {

	providerID := bson.ObjectIdHex(uid)
	return providerDAO.Update(providerID, provider)

}
