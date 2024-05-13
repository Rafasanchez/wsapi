package models

import (
	"time"

	"gorm.io/gorm"
)

type Proveedor struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `gorm:"not null"`
	Direccion string         `gorm:"not null"`
	Telefono  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateProveedor creates a new Proveedor record in the database.
func CreateProveedor(db *gorm.DB, proveedor *Proveedor) error {
	return db.Create(proveedor).Error
}

// GetProveedorByID retrieves a Proveedor record from the database by its ID.
func GetProveedorByID(db *gorm.DB, id uint) (*Proveedor, error) {
	var proveedor Proveedor
	err := db.First(&proveedor, id).Error
	if err != nil {
		return nil, err
	}
	return &proveedor, nil
}

// UpdateProveedor updates an existing Proveedor record in the database.
func UpdateProveedor(db *gorm.DB, proveedor *Proveedor) error {
	return db.Save(proveedor).Error
}

// DeleteProveedor deletes a Proveedor record from the database.
func DeleteProveedor(db *gorm.DB, proveedor *Proveedor) error {
	return db.Delete(proveedor).Error
}

// GetAllProveedores retrieves all Proveedor records from the database.
func GetAllProveedores(db *gorm.DB) ([]Proveedor, error) {
	var proveedores []Proveedor
	err := db.Find(&proveedores).Error
	if err != nil {
		return nil, err
	}
	return proveedores, nil
}