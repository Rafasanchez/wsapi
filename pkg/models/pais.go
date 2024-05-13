package models

import (
	"time"

	"gorm.io/gorm"
)

type Pais struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `gorm:"not null"`
	MonedaID  uint           `gorm:"not null"`
	Moneda    Moneda         `gorm:"foreignKey:MonedaID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}