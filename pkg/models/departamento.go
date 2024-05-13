package models

import (
	"time"

	"gorm.io/gorm"
)

type Departamento struct {
	ID        uint           `gorm:"primaryKey"`
	Nombre    string         `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PaisID    uint
	Pais      Pais
}

func CreateDepartamento(db *gorm.DB, departamento *Departamento) error {
	return db.Create(departamento).Error
}

func GetDepartamento(db *gorm.DB, id uint) (*Departamento, error) {
	var departamento Departamento
	err := db.Preload("Pais").First(&departamento, id).Error
	if err != nil {
		return nil, err
	}
	return &departamento, nil
}

func UpdateDepartamento(db *gorm.DB, departamento *Departamento) error {
	return db.Save(departamento).Error
}

func DeleteDepartamento(db *gorm.DB, id uint) error {
	return db.Delete(&Departamento{}, id).Error
}

func GetAllDepartamentos(db *gorm.DB) ([]Departamento, error) {
	var departamentos []Departamento
	err := db.Preload("Pais").Find(&departamentos).Error
	if err != nil {
		return nil, err
	}
	return departamentos, nil
}