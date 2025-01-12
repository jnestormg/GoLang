package models

import (
	"time"

	"gorm.io/gorm"
)

type Estado string

const (
	Pendiente Estado = "pendiente"
	Realizada Estado = "realizada"
	Cancelada Estado = "cancelada"
)
type Citas struct {
	gorm.Model
	Fecha time.Time `gorm:"type:date" json:"fecha"`
	Motivo string `json:"motivo"`
	Estado Estado `json:"estado"`
	PacienteID uint `json:"paciente_id"`
	MedicoID uint `json:"medico_id"`
}