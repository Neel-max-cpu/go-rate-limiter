package handlers

import (
	"net/http"

	"github.com/Neel-max-cpu/go-rate-limiter/internal/db"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/models"
	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	apiKey := c.GetHeader("x-api-key") // Get it again, or better yet:

	// query := `SELECT * FROM users WHERE api_key=$1` -- both works just below one controlled
	query := `SELECT id, email, api_key, created_at FROM users WHERE api_key=$1`

	var user models.User
	// or -- user := models.User{}

	err := db.DB.QueryRow(c, query, apiKey).Scan(
		&user.ID,
		&user.Email,
		&user.APIKey,
		&user.CreatedAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
