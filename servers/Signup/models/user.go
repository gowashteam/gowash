package models

import (
	"time"

	"github.com/okeyonyia123/gowash/servers/signup/util"
)

type User struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Username          string `json:"username" bson:"username"`
	FirstName         string `json:"firstName" bson:"firstName"`
	LastName          string `json:"lastName" bson:"lastName"`
	Email             string `json:"email" bson:"email"`
	PasswordHash      string `json:"passwordHash" bson:"passwordHash"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModified" bson:"timestampModified"`
}

func NewUser(FieldValues map[string]string) *User {

	passwordHash := util.SHA256OfString(FieldValues["password"])
	now := time.Now()
	unixTimestamp := now.Unix()
	u := User{
		UUID:              util.GenerateUUID(),
		Username:          FieldValues["username"],
		FirstName:         FieldValues["firstName"],
		LastName:          FieldValues["lastName"],
		Email:             FieldValues["email"],
		PasswordHash:      passwordHash,
		TimestampCreated:  unixTimestamp,
		TimestampModified: unixTimestamp,
	}
	return &u
}
