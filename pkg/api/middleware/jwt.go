package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware is a middleware function that handles JWT authentication for the API routes.
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// TODO: Replace "YOUR_SECRET_KEY" with your actual secret key used for signing the JWT tokens
			return []byte("YOUR_SECRET_KEY"), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// TODO: You can access the claims data like claims["userId"], claims["role"], etc.
		// and perform any necessary authorization checks here

		c.Next()
	}
}