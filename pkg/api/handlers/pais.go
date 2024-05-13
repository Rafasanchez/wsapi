package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetPaisByID retrieves a Pais by its ID
func GetPaisByID(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to retrieve Pais from the database by ID

	// Example response
	pais := models.Pais{
		ID:   id,
		Name: "Example Pais",
	}

	c.JSON(http.StatusOK, pais)
}

// CreatePais creates a new Pais
func CreatePais(c *gin.Context) {
	var pais models.Pais

	if err := c.ShouldBindJSON(&pais); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement code to create Pais in the database

	c.JSON(http.StatusCreated, pais)
}

// UpdatePais updates an existing Pais
func UpdatePais(c *gin.Context) {
	id := c.Param("id")
	var pais models.Pais

	if err := c.ShouldBindJSON(&pais); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement code to update Pais in the database by ID

	c.JSON(http.StatusOK, pais)
}

// DeletePais deletes a Pais by its ID
func DeletePais(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to delete Pais from the database by ID

	c.JSON(http.StatusOK, gin.H{"message": "Pais deleted"})
}
