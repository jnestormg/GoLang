package models

import "gorm.io/gorm"

type Medicos struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
	Cedula string `json:"cedula"`
	Domicilio Domicilio `gorm:"embedded" json:"domicilio"`
	Especialidades Especialidades `gorm:"foreignKey:medico_id" json:"especialidades"`
	Citas []Citas `gorm:"foreignKey:medico_id" json:"citas"`
}