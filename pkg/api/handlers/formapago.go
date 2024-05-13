package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetFormaPagoByID retrieves a FormaPago by ID
func GetFormaPagoByID(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to retrieve FormaPago from the database by ID

	// Example response
	formaPago := models.FormaPago{
		ID:   id,
		Name: "Example FormaPago",
	}

	c.JSON(http.StatusOK, formaPago)
}

// CreateFormaPago creates a new FormaPago
func CreateFormaPago(c *gin.Context) {
	var formaPago models.FormaPago

	// TODO: Implement code to parse request body and create FormaPago in the database

	// Example response
	formaPago.ID = "1"
	formaPago.Name = "Example FormaPago"

	c.JSON(http.StatusCreated, formaPago)
}

// UpdateFormaPago updates an existing FormaPago
func UpdateFormaPago(c *gin.Context) {
	id := c.Param("id")
	var formaPago models.FormaPago

	// TODO: Implement code to parse request body and update FormaPago in the database

	// Example response
	formaPago.ID = id
	formaPago.Name = "Updated FormaPago"

	c.JSON(http.StatusOK, formaPago)
}

// DeleteFormaPago deletes a FormaPago
func DeleteFormaPago(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to delete FormaPago from the database by ID

	c.JSON(http.StatusNoContent, nil)
}
