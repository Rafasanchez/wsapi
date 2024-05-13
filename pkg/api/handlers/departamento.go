package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetDepartamentoByID retrieves a departamento by its ID
func GetDepartamentoByID(c *gin.Context) {
	// Get the departamento ID from the request parameters
	departamentoID := c.Param("id")

	// Query the database for the departamento with the given ID
	var departamento models.Departamento
	if err := db.First(&departamento, departamentoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Departamento not found"})
		return
	}

	// Return the departamento as JSON response
	c.JSON(http.StatusOK, departamento)
}

// CreateDepartamento creates a new departamento
func CreateDepartamento(c *gin.Context) {
	// Bind the JSON request body to the departamento struct
	var departamento models.Departamento
	if err := c.ShouldBindJSON(&departamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the departamento in the database
	if err := db.Create(&departamento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create departamento"})
		return
	}

	// Return the created departamento as JSON response
	c.JSON(http.StatusCreated, departamento)
}

// UpdateDepartamento updates an existing departamento
func UpdateDepartamento(c *gin.Context) {
	// Get the departamento ID from the request parameters
	departamentoID := c.Param("id")

	// Query the database for the departamento with the given ID
	var departamento models.Departamento
	if err := db.First(&departamento, departamentoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Departamento not found"})
		return
	}

	// Bind the JSON request body to the departamento struct
	if err := c.ShouldBindJSON(&departamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the departamento in the database
	if err := db.Save(&departamento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update departamento"})
		return
	}

	// Return the updated departamento as JSON response
	c.JSON(http.StatusOK, departamento)
}

// DeleteDepartamento deletes a departamento
func DeleteDepartamento(c *gin.Context) {
	// Get the departamento ID from the request parameters
	departamentoID := c.Param("id")

	// Query the database for the departamento with the given ID
	var departamento models.Departamento
	if err := db.First(&departamento, departamentoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Departamento not found"})
		return
	}

	// Delete the departamento from the database
	if err := db.Delete(&departamento).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete departamento"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Departamento deleted successfully"})
}
