package models

import (
	"time"

	"gorm.io/gorm"
)

type Municipio struct {
	ID           uint           `gorm:"primaryKey"`
	Nombre       string         `gorm:"not null"`
	Departamento Departamento   `gorm:"foreignKey:DepartamentoID"`
	DepartamentoID uint         `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}