package models

import "gorm.io/gorm"

type Especialidades struct {
	gorm.Model
	Nombre string
	Descripcion string
	UserID uint
}