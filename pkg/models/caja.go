package models

import (
	"time"

	"gorm.io/gorm"
)

type Caja struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `gorm:"not null"`
	FechaAp   time.Time      `gorm:"not null"`
	FechaCie  *time.Time
	Saldo     float64        `gorm:"not null;default:0"`
	EmpresaID uint           `gorm:"not null"`
	Empresa   Empresa        `gorm:"foreignKey:EmpresaID"`
}

func CreateCaja(db *gorm.DB, caja *Caja) error {
	return db.Create(caja).Error
}

func GetCaja(db *gorm.DB, id uint) (*Caja, error) {
	var caja Caja
	err := db.Preload("Empresa").First(&caja, id).Error
	if err != nil {
		return nil, err
	}
	return &caja, nil
}

func UpdateCaja(db *gorm.DB, caja *Caja) error {
	return db.Save(caja).Error
}

func DeleteCaja(db *gorm.DB, id uint) error {
	return db.Delete(&Caja{}, id).Error
}

func GetAllCajas(db *gorm.DB) ([]Caja, error) {
	var cajas []Caja
	err := db.Preload("Empresa").Find(&cajas).Error
	if err != nil {
		return nil, err
	}
	return cajas, nil
}