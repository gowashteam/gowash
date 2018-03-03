package datastore

import (
	"errors"

	"github.com/okeyonyia123/gowash/servers/Signup/models"
)

type Datastore interface {
	PostUser(user *models.User) (*models.User, error)
	GetUser(username string) (*models.User, error)
	Close()
}

const (
	MONGODB = iota
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	default:
		return nil, errors.New("The datastore you specified does not exist!")
	}

}
