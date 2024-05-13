package models

import (
	"time"

	"gorm.io/gorm"
)

type Producto struct {
	ID          uint           `gorm:"primaryKey"`
	Nombre      string         `gorm:"not null"`
	Precio      float64        `gorm:"not null"`
	Descripcion string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// CreateProducto creates a new producto record in the database.
func CreateProducto(db *gorm.DB, producto *Producto) error {
	return db.Create(producto).Error
}

// GetProductoByID retrieves a producto record from the database by its ID.
func GetProductoByID(db *gorm.DB, id uint) (*Producto, error) {
	var producto Producto
	err := db.First(&producto, id).Error
	if err != nil {
		return nil, err
	}
	return &producto, nil
}

// UpdateProducto updates an existing producto record in the database.
func UpdateProducto(db *gorm.DB, producto *Producto) error {
	return db.Save(producto).Error
}

// DeleteProducto deletes a producto record from the database.
func DeleteProducto(db *gorm.DB, producto *Producto) error {
	return db.Delete(producto).Error
}

// GetAllProductos retrieves all producto records from the database.
func GetAllProductos(db *gorm.DB) ([]Producto, error) {
	var productos []Producto
	err := db.Find(&productos).Error
	if err != nil {
		return nil, err
	}
	return productos, nil
}