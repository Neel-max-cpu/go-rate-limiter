package middleware

import (
	"net/http"

	"github.com/Neel-max-cpu/go-rate-limiter/internal/db"
	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.GetHeader("x-api-key")

		if apikey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing api key",
			})
			c.Abort()
			return
		}

		var exists bool

		query := `SELECT EXISTS(SELECT 1 FROM users WHERE api_key=$1)`
		err := db.DB.QueryRow(c, query, apikey).Scan(&exists)

		if err != nil || !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid api key!",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
