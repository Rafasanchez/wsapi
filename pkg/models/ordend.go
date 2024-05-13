package models

import (
	"time"

	"gorm.io/gorm"
)

type Ordend struct {
	ID          uint           `gorm:"primaryKey"`
	OrdenhID    uint           `gorm:"not null"`
	ProductoID  uint           `gorm:"not null"`
	Cantidad    int            `gorm:"not null"`
	Precio      float64        `gorm:"not null"`
	Total       float64        `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func CreateOrdend(ordend *Ordend) error {
	return db.Create(ordend).Error
}

func GetOrdendByID(id uint) (*Ordend, error) {
	var ordend Ordend
	err := db.First(&ordend, id).Error
	if err != nil {
		return nil, err
	}
	return &ordend, nil
}

func UpdateOrdend(ordend *Ordend) error {
	return db.Save(ordend).Error
}

func DeleteOrdend(ordend *Ordend) error {
	return db.Delete(ordend).Error
}

func GetAllOrdends() ([]Ordend, error) {
	var ordends []Ordend
	err := db.Find(&ordends).Error
	if err != nil {
		return nil, err
	}
	return ordends, nil
}