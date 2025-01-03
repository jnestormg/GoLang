package models

import "gorm.io/gorm"

type Pacientes struct{
	gorm.Model
	Nombre string
	Apellido string
	Edad int
	Genero string
	Telefono string
	Domicilio Domicilio `gorm:"embedded"`
}