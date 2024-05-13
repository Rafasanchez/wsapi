package handlers

import (
	"net/http"

	"github.com/Rafasanchez/wsapi/pkg/models"
	"github.com/gin-gonic/gin"
)

// GetListaPreciohByID handles the GET request to retrieve a ListaPrecioh by ID
func GetListaPreciohByID(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to fetch ListaPrecioh from the database by ID

	// Example response
	listaPrecioh := models.ListaPrecioh{
		ID:          id,
		Nombre:      "Lista de Precios 1",
		Descripcion: "Lista de precios para productos",
	}

	c.JSON(http.StatusOK, listaPrecioh)
}

// CreateListaPrecioh handles the POST request to create a new ListaPrecioh
func CreateListaPrecioh(c *gin.Context) {
	var listaPrecioh models.ListaPrecioh

	if err := c.ShouldBindJSON(&listaPrecioh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement code to create a new ListaPrecioh in the database

	// Example response
	listaPrecioh.ID = "1"

	c.JSON(http.StatusCreated, listaPrecioh)
}

// UpdateListaPrecioh handles the PUT request to update an existing ListaPrecioh
func UpdateListaPrecioh(c *gin.Context) {
	id := c.Param("id")
	var listaPrecioh models.ListaPrecioh

	if err := c.ShouldBindJSON(&listaPrecioh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement code to update the ListaPrecioh in the database by ID

	// Example response
	listaPrecioh.ID = id

	c.JSON(http.StatusOK, listaPrecioh)
}

// DeleteListaPrecioh handles the DELETE request to delete a ListaPrecioh by ID
func DeleteListaPrecioh(c *gin.Context) {
	id := c.Param("id")

	// TODO: Implement code to delete the ListaPrecioh from the database by ID

	c.JSON(http.StatusNoContent, nil)
}
