package models

import (
	"time"
)

// Empleado represents the Empleado table in the database
type Empleado struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Nombre      string     `gorm:"not null" json:"nombre"`
	Apellido    string     `gorm:"not null" json:"apellido"`
	Email       string     `gorm:"not null;unique" json:"email"`
	FechaNac    *time.Time `gorm:"type:date" json:"fecha_nac"`
	Direccion   string     `json:"direccion"`
	Ciudad      string     `json:"ciudad"`
	Departamento string    `json:"departamento"`
	Pais        string     `json:"pais"`
	SucursalID  uint       `gorm:"not null" json:"sucursal_id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
}

// CreateEmpleado creates a new Empleado record in the database
func CreateEmpleado(empleado *Empleado) error {
	return db.Create(empleado).Error
}

// GetEmpleadoByID retrieves an Empleado record from the database by ID
func GetEmpleadoByID(id uint) (*Empleado, error) {
	var empleado Empleado
	err := db.First(&empleado, id).Error
	if err != nil {
		return nil, err
	}
	return &empleado, nil
}

// UpdateEmpleado updates an existing Empleado record in the database
func UpdateEmpleado(empleado *Empleado) error {
	return db.Save(empleado).Error
}

// DeleteEmpleado deletes an Empleado record from the database
func DeleteEmpleado(empleado *Empleado) error {
	return db.Delete(empleado).Error
}

// GetAllEmpleados retrieves all Empleado records from the database
func GetAllEmpleados() ([]Empleado, error) {
	var empleados []Empleado
	err := db.Find(&empleados).Error
	if err != nil {
		return nil, err
	}
	return empleados, nil
}