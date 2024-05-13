package models

import (
	"time"

	"gorm.io/gorm"
)

type ListaPreciod struct {
	ID           uint           `gorm:"primaryKey"`
	ListaPrecioh ListaPrecioh   `gorm:"foreignKey:ListaPreciohID"`
	Producto     Producto       `gorm:"foreignKey:ProductoID"`
	Precio       float64        `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// CreateListaPreciod creates a new ListaPreciod record in the database.
func CreateListaPreciod(listaPreciod *ListaPreciod) error {
	return db.Create(listaPreciod).Error
}

// GetListaPreciodByID retrieves a ListaPreciod record from the database by its ID.
func GetListaPreciodByID(id uint) (*ListaPreciod, error) {
	var listaPreciod ListaPreciod
	err := db.Preload("ListaPrecioh").Preload("Producto").First(&listaPreciod, id).Error
	if err != nil {
		return nil, err
	}
	return &listaPreciod, nil
}

// UpdateListaPreciod updates an existing ListaPreciod record in the database.
func UpdateListaPreciod(listaPreciod *ListaPreciod) error {
	return db.Save(listaPreciod).Error
}

// DeleteListaPreciod deletes a ListaPreciod record from the database.
func DeleteListaPreciod(listaPreciod *ListaPreciod) error {
	return db.Delete(listaPreciod).Error
}