package routes

import (
	"time"

	"github.com/Neel-max-cpu/go-rate-limiter/internal/handlers"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		/*
			func(c *gin.Context) = This is an anonymous function, also known as the handler function for this route. Gin will execute this code whenever a GET request to /health is received. The c *gin.Context parameter is a powerful object that carries the request details and provides methods to send a response.
		*/
		c.JSON(200, gin.H{
			"message": "server running",
		})
	})

	r.POST("/users", handlers.CreateUser)

	protected := r.Group("/")
	protected.Use(
		middleware.APIKeyAuth(),
		middleware.RateLimit(10, time.Minute), // 10 requests per minute
	)

	protected.GET("/data", handlers.GetData)
}
