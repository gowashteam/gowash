package datastore

import (
	"errors"

	"github.com/okeyonyia123/gowash/servers/signup/models"
)

type Datastore interface {
	CreateUser(user *models.User) error
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