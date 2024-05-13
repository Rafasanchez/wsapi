package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrdendByID handles the GET request to retrieve an Ordend by ID.
func GetOrdendByID(c *gin.Context) {
	// TODO: Implement logic to retrieve an Ordend by ID from the database.
	ordendID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "GetOrdendByID", "id": ordendID})
}

// CreateOrdend handles the POST request to create a new Ordend.
func CreateOrdend(c *gin.Context) {
	// TODO: Implement logic to create a new Ordend in the database.
	var ordend Ordend
	if err := c.ShouldBindJSON(&ordend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CreateOrdend", "ordend": ordend})
}

// UpdateOrdend handles the PUT request to update an existing Ordend.
func UpdateOrdend(c *gin.Context) {
	// TODO: Implement logic to update an existing Ordend in the database.
	ordendID := c.Param("id")
	var updatedOrdend Ordend
	if err := c.ShouldBindJSON(&updatedOrdend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "UpdateOrdend", "id": ordendID, "ordend": updatedOrdend})
}

// DeleteOrdend handles the DELETE request to delete an Ordend.
func DeleteOrdend(c *gin.Context) {
	// TODO: Implement logic to delete an Ordend from the database.
	ordendID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "DeleteOrdend", "id": ordendID})
}