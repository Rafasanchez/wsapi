package routes

import (
	"github.com/Rafasanchez/wsapi/pkg/api/handlers"
	"github.com/Rafasanchez/wsapi/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the API routes and associates them with the corresponding handler functions
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/login", handlers.Login)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/monedas", handlers.GetMonedas)
		protected.POST("/monedas", handlers.CreateMoneda)
		protected.GET("/monedas/:id", handlers.GetMonedaByID)
		protected.PUT("/monedas/:id", handlers.UpdateMoneda)
		protected.DELETE("/monedas/:id", handlers.DeleteMoneda)

		protected.GET("/paises", handlers.GetPaises)
		protected.POST("/paises", handlers.CreatePais)
		protected.GET("/paises/:id", handlers.GetPaisByID)
		protected.PUT("/paises/:id", handlers.UpdatePais)
		protected.DELETE("/paises/:id", handlers.DeletePais)

		// Add routes for other tables here
	}

	return router
}
