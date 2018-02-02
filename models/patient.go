package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Patient estructura de la tabla paciente
type Patient struct {
	ID           bson.ObjectId `bson:"_id, omitempty"`
	DateInit     time.Time     `bson:", omitempty"`
	NomApe       string        `bson:", omitempty"`
	Age          uint          `bson:", omitempty"`
	Sex          string        `bson:", omitempty"`
	DateNac      time.Time     `bson:", omitempty"`
	Address      string
	Ocupation    string `bson:", omitempty"`
	TelCel       string
	Alergies     string
	Operations   string
	Diabettes    bool `bson:", omitempty"`
	Hipertension bool `bson:", omitempty"`
	Others       string
	TreatMedics  string
	State        bool `bson:", omitempty"`
	TreatmentID  uint
	CreateAt     time.Time `bson:", omitempty"`
	UpdateAt     time.Time
	DeleteAt     time.Time
}

// Patients slice de pacientes
type Patients []Patient
