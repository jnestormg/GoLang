package models

import "gorm.io/gorm"

type Pacientes struct{
	gorm.Model
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad int `json:"edad"`
	Genero string `json:"genero"`
	Telefono string `json:"telefono"`
	Domicilio Domicilio `gorm:"embedded" json:"domicilio"`
	Citas []Citas `gorm:"foreignKey:paciente_id" json:"citas"`
}