package daos

import (
	"github.com/mperezguendulain/el-gallero-api-rest/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ProviderDAO is the DAO of Provider
type ProviderDAO struct {
	Collection *mgo.Collection
}

// FindAll get all providers in Providers Collection
func (pd *ProviderDAO) FindAll() ([]models.Provider, error) {

	var providers []models.Provider
	err := pd.Collection.Find(bson.M{}).All(&providers)
	return providers, err

}

// FindByID search for a provider by his id
func (pd *ProviderDAO) FindByID(providerID string) (models.Provider, error) {

	var provider models.Provider
	err := pd.Collection.FindId(bson.ObjectIdHex(providerID)).One(&provider)
	return provider, err

}

// Insert inserts a document in providers Collection
func (pd *ProviderDAO) Insert(provider models.Provider) error {

	err := pd.Collection.Insert(&provider)
	return err

}

// RemoveByID removes a provider in providers Collection
func (pd *ProviderDAO) RemoveByID(providerID bson.ObjectId) error {

	err := pd.Collection.RemoveId(providerID)
	return err

}

// Update updates a provider in providers Collection
func (pd *ProviderDAO) Update(providerID bson.ObjectId, provider models.Provider) error {

	err := pd.Collection.UpdateId(providerID, &provider)
	return err

}
