package models

import (
	"time"

	"gorm.io/gorm"
)

type Sucursal struct {
	ID          uint           `gorm:"primaryKey"`
	Nombre      string         `gorm:"not null"`
	Direccion   string         `gorm:"not null"`
	Ciudad      string         `gorm:"not null"`
	Telefono    string         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	EmpresaID   uint           `gorm:"not null"`
	Empresa     Empresa        `gorm:"foreignKey:EmpresaID"`
	Empleados   []Empleado     `gorm:"foreignKey:SucursalID"`
	Ordenes     []Ordenh       `gorm:"foreignKey:SucursalID"`
}

func CreateSucursal(db *gorm.DB, sucursal *Sucursal) error {
	return db.Create(sucursal).Error
}

func GetSucursal(db *gorm.DB, id uint) (*Sucursal, error) {
	var sucursal Sucursal
	err := db.Preload("Empresa").Preload("Empleados").Preload("Ordenes").First(&sucursal, id).Error
	if err != nil {
		return nil, err
	}
	return &sucursal, nil
}

func UpdateSucursal(db *gorm.DB, sucursal *Sucursal) error {
	return db.Save(sucursal).Error
}

func DeleteSucursal(db *gorm.DB, id uint) error {
	return db.Delete(&Sucursal{}, id).Error
}

func GetAllSucursales(db *gorm.DB) ([]Sucursal, error) {
	var sucursales []Sucursal
	err := db.Preload("Empresa").Preload("Empleados").Preload("Ordenes").Find(&sucursales).Error
	if err != nil {
		return nil, err
	}
	return sucursales, nil
}