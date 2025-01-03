package models

import "gorm.io/gorm"

type Especialidades struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`	
	UserID uint `gorm:"column:user_id" json:"user_id"`
}