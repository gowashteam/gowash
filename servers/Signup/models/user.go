package models

import (
	"time"

	"github.com/okeyonyia123/gowash/servers/Signup/handlers"
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

func NewUser(form *handlers.Form) *User {

	passwordHash := util.SHA256OfString(form.FieldValues[password])
	now := time.Now()
	unixTimestamp := now.Unix()
	u := User{
		UUID: util.GenerateUUID(),
		Username:          form.FieldValues[username],
		FirstName:         form.FieldValues[firstName],
		LastName:          form.FieldValues[lastName],
		Email:             form.FieldValues[email],
		PasswordHash:      passwordHash,
		TimestampCreated:  unixTimestamp,
		TimestampModified: unixTimestamp
	}
	return &u
}
