package models

import (
	"time"

	"gorm.io/gorm"
)

type Empresa struct {
	ID          uint           `gorm:"primaryKey"`
	Nombre      string         `gorm:"not null"`
	Direccion   string         `gorm:"not null"`
	Telefono    string         `gorm:"not null"`
	FechaCreacion time.Time     `gorm:"not null"`
	PaisID      uint           `gorm:"not null"`
	Pais        Pais           `gorm:"foreignKey:PaisID"`
	Clientes    []Cliente      `gorm:"foreignKey:EmpresaID"`
	Proveedores []Proveedor    `gorm:"foreignKey:EmpresaID"`
	Productos   []Producto     `gorm:"foreignKey:EmpresaID"`
	Cajas       []Caja         `gorm:"foreignKey:EmpresaID"`
	ListaPrecios []ListaPrecioh `gorm:"foreignKey:EmpresaID"`
	Sucursales  []Sucursal     `gorm:"foreignKey:EmpresaID"`
}

func CreateEmpresa(db *gorm.DB, empresa *Empresa) error {
	return db.Create(empresa).Error
}

func GetEmpresa(db *gorm.DB, id uint) (*Empresa, error) {
	var empresa Empresa
	err := db.Preload("Pais").Preload("Clientes").Preload("Proveedores").Preload("Productos").Preload("Cajas").Preload("ListaPrecios").Preload("Sucursales").First(&empresa, id).Error
	if err != nil {
		return nil, err
	}
	return &empresa, nil
}

func UpdateEmpresa(db *gorm.DB, empresa *Empresa) error {
	return db.Save(empresa).Error
}

func DeleteEmpresa(db *gorm.DB, id uint) error {
	return db.Delete(&Empresa{}, id).Error
}

func GetAllEmpresas(db *gorm.DB) ([]Empresa, error) {
	var empresas []Empresa
	err := db.Preload("Pais").Preload("Clientes").Preload("Proveedores").Preload("Productos").Preload("Cajas").Preload("ListaPrecios").Preload("Sucursales").Find(&empresas).Error
	if err != nil {
		return nil, err
	}
	return empresas, nil
}