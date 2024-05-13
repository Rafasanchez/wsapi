package models

import (
	"time"

	"gorm.io/gorm"
)

type FormaPago struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateFormaPago creates a new FormaPago record in the database.
func CreateFormaPago(db *gorm.DB, formaPago *FormaPago) error {
	return db.Create(formaPago).Error
}

// GetFormaPagoByID retrieves a FormaPago record from the database by ID.
func GetFormaPagoByID(db *gorm.DB, id uint) (*FormaPago, error) {
	var formaPago FormaPago
	err := db.First(&formaPago, id).Error
	if err != nil {
		return nil, err
	}
	return &formaPago, nil
}

// UpdateFormaPago updates an existing FormaPago record in the database.
func UpdateFormaPago(db *gorm.DB, formaPago *FormaPago) error {
	return db.Save(formaPago).Error
}

// DeleteFormaPago deletes a FormaPago record from the database.
func DeleteFormaPago(db *gorm.DB, formaPago *FormaPago) error {
	return db.Delete(formaPago).Error
}

// GetAllFormaPagos retrieves all FormaPago records from the database.
func GetAllFormaPagos(db *gorm.DB) ([]FormaPago, error) {
	var formaPagos []FormaPago
	err := db.Find(&formaPagos).Error
	if err != nil {
		return nil, err
	}
	return formaPagos, nil
}