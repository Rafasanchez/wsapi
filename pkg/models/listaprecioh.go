package models

import (
	"time"

	"gorm.io/gorm"
)

type ListaPrecioh struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `gorm:"not null"`
	Fecha     time.Time      `gorm:"not null"`
	EmpresaID uint           `gorm:"not null"`
	Empresa   Empresa        `gorm:"foreignKey:EmpresaID"`
	ListaPreciods []ListaPreciod `gorm:"foreignKey:ListaPreciohID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}