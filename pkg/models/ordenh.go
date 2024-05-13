package models

import (
	"time"
)

type Ordenh struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmpresaID   uint       `gorm:"not null" json:"empresa_id"`
	SucursalID  uint       `gorm:"not null" json:"sucursal_id"`
	FormaPagoID uint       `gorm:"not null" json:"forma_pago_id"`
	Numero      string     `gorm:"not null" json:"numero"`
	Fecha       time.Time  `gorm:"not null" json:"fecha"`
	Total       float64    `gorm:"not null" json:"total"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `gorm:"index" json:"-"`
}

// CreateOrdenh creates a new Ordenh record in the database.
func CreateOrdenh(ordenh *Ordenh) error {
	return db.Create(ordenh).Error
}

// GetOrdenhByID retrieves an Ordenh record from the database by its ID.
func GetOrdenhByID(id uint) (*Ordenh, error) {
	var ordenh Ordenh
	err := db.First(&ordenh, id).Error
	if err != nil {
		return nil, err
	}
	return &ordenh, nil
}

// UpdateOrdenh updates an existing Ordenh record in the database.
func UpdateOrdenh(ordenh *Ordenh) error {
	return db.Save(ordenh).Error
}

// DeleteOrdenh deletes an Ordenh record from the database.
func DeleteOrdenh(ordenh *Ordenh) error {
	return db.Delete(ordenh).Error
}

// GetAllOrdenhs retrieves all Ordenh records from the database.
func GetAllOrdenhs() ([]Ordenh, error) {
	var ordenhs []Ordenh
	err := db.Find(&ordenhs).Error
	if err != nil {
		return nil, err
	}
	return ordenhs, nil
}