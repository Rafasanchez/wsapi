package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetCajaByID retrieves a Caja by its ID
func GetCajaByID(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to retrieve Caja from the database by ID

	// Example response
	caja := models.Caja{
		ID:          id,
		Description: "Example Caja",
	}

	c.JSON(http.StatusOK, caja)
}

// CreateCaja creates a new Caja
func CreateCaja(c *gin.Context) {
	var caja models.Caja

	// TODO: Implement code to parse request body and create Caja in the database

	// Example response
	caja.ID = "1"
	caja.Description = "New Caja"

	c.JSON(http.StatusCreated, caja)
}

// UpdateCaja updates an existing Caja
func UpdateCaja(c *gin.Context) {
	id := c.Param("id")
	var caja models.Caja

	// TODO: Implement code to parse request body and update Caja in the database

	// Example response
	caja.ID = id
	caja.Description = "Updated Caja"

	c.JSON(http.StatusOK, caja)
}

// DeleteCaja deletes a Caja by its ID
func DeleteCaja(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to delete Caja from the database by ID

	c.JSON(http.StatusNoContent, nil)
}
