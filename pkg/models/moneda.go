package models

import (
	"time"

	"gorm.io/gorm"
)

// Moneda represents the Moneda table in the database
type Moneda struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `gorm:"not null"`
	Simbolo   string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}