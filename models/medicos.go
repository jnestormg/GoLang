package models

import "gorm.io/gorm"

type Medicos struct {
	gorm.Model
	ID uint `gorm:"primaryKey" json:"id"`
	Nombre string
	Apellido string
	Telefono string
	Cedula string
	Domicilio Domicilio `gorm:"embedded"`
	Especialidades Especialidades 
	Citas []Citas
}