package models

import (
	"time"

	"gorm.io/gorm"
)

type Estado string

const (
	Activo Estado = "Activo"
	Inactivo Estado = "Inactivo"
)
type Citas struct {
	gorm.Model
	Fecha time.Time `gorm:"type:date" json:"fecha"`
	Motivo string `json:"motivo"`
	Estado Estado `json:"estado"`
	MedicoID uint `json:"medico_id"`	
	PacienteID uint `json:"paciente_id"`
}