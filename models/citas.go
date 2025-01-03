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
	MedicoID uint `gorm:"column:medico_id" json:"medico_id"`	
	PacienteID uint `gorm:"column:paciente_id" json:"paciente_id"`
}