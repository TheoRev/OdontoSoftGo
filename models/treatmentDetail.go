package models

import "gopkg.in/mgo.v2/bson"

// TreatmentDetail estructura q contienes todas las curaciones realizadas al paciente
type TreatmentDetail struct {
	ID          bson.ObjectId `bson:"_id, omitempty"`
	WorkID      uint          `bson:", omitempty"`
	Quantity    uint          `bson:", omitempty"`
	TreatmentID uint          `bson:", omitempty"`
}

// TreatmentsDetail slice de curaciones de los tratamientos
type TreatmentsDetail []TreatmentDetail
