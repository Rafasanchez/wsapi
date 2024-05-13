package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/api/models"
	"github.com/gin-gonic/gin"
)

// GetEmpresaByID retrieves an Empresa by ID
func GetEmpresaByID(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement logic to retrieve Empresa from the database by ID

	// Example response
	empresa := models.Empresa{
		ID:   id,
		Name: "Example Empresa",
	}

	c.JSON(http.StatusOK, empresa)
}

// CreateEmpresa creates a new Empresa
func CreateEmpresa(c *gin.Context) {
	var empresa models.Empresa

	// TODO: Implement logic to create a new Empresa in the database

	// Example request body
	if err := c.ShouldBindJSON(&empresa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example response
	empresa.ID = "1"

	c.JSON(http.StatusCreated, empresa)
}

// UpdateEmpresa updates an existing Empresa
func UpdateEmpresa(c *gin.Context) {
	id := c.Param("id")
	var empresa models.Empresa

	// TODO: Implement logic to update the existing Empresa in the database

	// Example request body
	if err := c.ShouldBindJSON(&empresa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example response
	empresa.ID = id

	c.JSON(http.StatusOK, empresa)
}

// DeleteEmpresa deletes an existing Empresa
func DeleteEmpresa(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement logic to delete the existing Empresa from the database

	c.JSON(http.StatusNoContent, nil)
}
