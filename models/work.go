package models

import "gopkg.in/mgo.v2/bson"

// Work estructura de la tabla trabajo
type Work struct {
	ID                bson.ObjectId `bson:"_id, omitempty"`
	Name              string        `bson:", omitempty"`
	Price             float64       `bson:", omitempty"`
	TreatmentDetailID uint          `bson:", omitempty"`
}

// Works slice de trabajos
type Works []Work
