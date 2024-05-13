package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetAllMonedas handles the GET request to retrieve all Monedas
func GetAllMonedas(c *gin.Context) {
	var monedas []models.Moneda
	err := models.DB.Find(&monedas).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, monedas)
}

// GetMonedaByID handles the GET request to retrieve a Moneda by ID
func GetMonedaByID(c *gin.Context) {
	id := c.Param("id")
	var moneda models.Moneda
	err := models.DB.First(&moneda, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Moneda not found"})
		return
	}
	c.JSON(http.StatusOK, moneda)
}

// CreateMoneda handles the POST request to create a new Moneda
func CreateMoneda(c *gin.Context) {
	var moneda models.Moneda
	err := c.ShouldBindJSON(&moneda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DB.Create(&moneda).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, moneda)
}

// UpdateMoneda handles the PUT request to update an existing Moneda
func UpdateMoneda(c *gin.Context) {
	id := c.Param("id")
	var moneda models.Moneda
	err := models.DB.First(&moneda, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Moneda not found"})
		return
	}
	err = c.ShouldBindJSON(&moneda)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.DB.Save(&moneda).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, moneda)
}

// DeleteMoneda handles the DELETE request to delete a Moneda
func DeleteMoneda(c *gin.Context) {
	id := c.Param("id")
	var moneda models.Moneda
	err := models.DB.First(&moneda, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Moneda not found"})
		return
	}
	err = models.DB.Delete(&moneda).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Moneda deleted successfully"})
}
