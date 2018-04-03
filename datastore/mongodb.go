package datastore

import (
	"github.com/okeyonyia123/gowash/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//class name with fields
type MongoDBDatastore struct {
	*mgo.Session
}

//constructor
func NewMongoDBDatastore(url string) (*MongoDBDatastore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDBDatastore{
		Session: session,
	}, nil
}

func (m *MongoDBDatastore) PostUser(user *models.User) (*models.User, error) {

	session := m.Copy()

	defer session.Close()
	userCollection := session.DB("gowash").C("User")
	err := userCollection.Insert(user)
	if err != nil {
		return nil, err
	}

	return m.GetUser(user.Username)
}

func (m *MongoDBDatastore) PostPartner(partner *models.Partner) (*models.Partner, error) {

	session := m.Copy()

	defer session.Close()
	PartnerCollection := session.DB("gowash").C("Partner")
	err := PartnerCollection.Insert(partner)
	if err != nil {
		return nil, err
	}

	return m.GetPartner(partner.Username)
}

func (m *MongoDBDatastore) GetUser(username string) (*models.User, error) {

	session := m.Copy()
	defer session.Close()
	userCollection := session.DB("gowash").C("User")
	u := models.User{}
	err := userCollection.Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (m *MongoDBDatastore) GetPartner(username string) (*models.Partner, error) {

	session := m.Copy()
	defer session.Close()
	userCollection := session.DB("gowash").C("Partner")
	u := models.Partner{}
	err := userCollection.Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
