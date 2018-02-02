package models

import "gopkg.in/mgo.v2/bson"

// Treatment estructura de la tabla tratamiento
type Treatment struct {
	ID                bson.ObjectId `bson:"_id, omitempty"`
	PatientID         uint          `bson:", omitempty"`
	TreatmentDetailID uint          `bson:", omitempty"`
	Select            bool          
}

// Treatments slice de tratamientos
type Treatments []Treatment
