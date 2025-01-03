package models

import "gorm.io/gorm"

type Medicos struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
	Cedula string `json:"cedula"`
	Domicilio Domicilio `gorm:"embedded" json:"domicilio"`
	Especialidades Especialidades `json:"especialidades"`
	Citas []Citas `json:"citas"`
}