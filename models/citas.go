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
	Fecha time.Time
	Motivo string
	Estado Estado
	CitasID uint
	MedicoID uint
}