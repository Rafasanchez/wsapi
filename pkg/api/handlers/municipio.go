package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetMunicipioByID retrieves a Municipio by its ID
func GetMunicipioByID(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to retrieve the Municipio from the database by ID

	// Example response
	municipio := models.Municipio{
		ID:   id,
		Name: "Example Municipio",
	}

	c.JSON(http.StatusOK, municipio)
}

// CreateMunicipio creates a new Municipio
func CreateMunicipio(c *gin.Context) {
	var municipio models.Municipio

	// TODO: Implement code to parse the request body and create a new Municipio in the database

	// Example response
	municipio.ID = "1"
	municipio.Name = "New Municipio"

	c.JSON(http.StatusCreated, municipio)
}

// UpdateMunicipio updates an existing Municipio
func UpdateMunicipio(c *gin.Context) {
	id := c.Param("id")
	var municipio models.Municipio

	// TODO: Implement code to parse the request body and update the Municipio in the database

	// Example response
	municipio.ID = id
	municipio.Name = "Updated Municipio"

	c.JSON(http.StatusOK, municipio)
}

// DeleteMunicipio deletes a Municipio by its ID
func DeleteMunicipio(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to delete the Municipio from the database by ID

	c.Status(http.StatusNoContent)
}
