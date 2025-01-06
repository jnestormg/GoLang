package models

import "gorm.io/gorm"

type Especialidades struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	MedicoID uint `json:"medico_id"`
}