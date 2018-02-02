package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User estructura de la tabla usuario
type User struct {
	ID       bson.ObjectId `bson:"_id, omitempty"`
	Username string        `bson:", omitempty"`
	Password string        `bson:", omitempty"`
	Level    string        `bson:", omitempty"`
}

// Users slice de usuarios
type Users []User
