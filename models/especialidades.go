package models

import "gorm.io/gorm"

type Especialidades struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`	
	UserID uint `json:"user_id"`
}