package models

import (
	"time"
)

// Cliente represents the Cliente model
type Cliente struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Nombre    string     `gorm:"not null" json:"nombre"`
	Email     string     `gorm:"not null;unique" json:"email"`
	Telefono  string     `gorm:"not null" json:"telefono"`
	Direccion string     `gorm:"not null" json:"direccion"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}

// CreateCliente creates a new Cliente record in the database
func CreateCliente(cliente *Cliente) error {
	return db.Create(cliente).Error
}

// GetClienteByID retrieves a Cliente record from the database by ID
func GetClienteByID(id uint) (*Cliente, error) {
	var cliente Cliente
	err := db.First(&cliente, id).Error
	return &cliente, err
}

// UpdateCliente updates an existing Cliente record in the database
func UpdateCliente(cliente *Cliente) error {
	return db.Save(cliente).Error
}

// DeleteCliente deletes a Cliente record from the database
func DeleteCliente(cliente *Cliente) error {
	return db.Delete(cliente).Error
}

// GetAllClientes retrieves all Cliente records from the database
func GetAllClientes() ([]Cliente, error) {
	var clientes []Cliente
	err := db.Find(&clientes).Error
	return clientes, err
}