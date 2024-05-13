package models

import (
	"time"
)

// CajaMov represents the CajaMov table in the database
type CajaMov struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CajaID      uint      `gorm:"not null" json:"caja_id"`
	TipoMov     string    `gorm:"not null" json:"tipo_mov"`
	Descripcion string    `gorm:"not null" json:"descripcion"`
	Monto       float64   `gorm:"not null" json:"monto"`
	Fecha       time.Time `gorm:"not null" json:"fecha"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// CreateCajaMov creates a new CajaMov record in the database
func CreateCajaMov(cajaMov *CajaMov) error {
	return db.Create(cajaMov).Error
}

// GetCajaMovByID retrieves a CajaMov record from the database by its ID
func GetCajaMovByID(id uint) (*CajaMov, error) {
	cajaMov := &CajaMov{}
	err := db.First(cajaMov, id).Error
	return cajaMov, err
}

// UpdateCajaMov updates an existing CajaMov record in the database
func UpdateCajaMov(cajaMov *CajaMov) error {
	return db.Save(cajaMov).Error
}

// DeleteCajaMov deletes a CajaMov record from the database
func DeleteCajaMov(cajaMov *CajaMov) error {
	return db.Delete(cajaMov).Error
}

// GetAllCajaMovs retrieves all CajaMov records from the database
func GetAllCajaMovs() ([]CajaMov, error) {
	var cajaMovs []CajaMov
	err := db.Find(&cajaMovs).Error
	return cajaMovs, err
}