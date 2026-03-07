package main

// This defines an executable program in Go. The main package is the entry point for the application.

import (
	"log"

	"github.com/Neel-max-cpu/go-rate-limiter/internal/db"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// This imports the Gin web framework, which provides all the tools for creating the web server and handling requests.

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// postgres connect
	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// redis connect
	redis.Connect()

	r := gin.Default()
	/*
		This creates a Gin router (or engine) with two essential middleware already attached:
		Logger: It automatically logs every HTTP request your server receives (useful for debugging and monitoring).
		Recovery: It automatically catches any panic (a critical error) that happens while handling a request, preventing your entire server from crashing. It will instead return a 500 Internal Server Error response.
	*/

	r.GET("/health", func(c *gin.Context) {
		/*
			func(c *gin.Context) = This is an anonymous function, also known as the handler function for this route. Gin will execute this code whenever a GET request to /health is received. The c *gin.Context parameter is a powerful object that carries the request details and provides methods to send a response.
		*/
		c.JSON(200, gin.H{
			"message": "server running",
		})
	})
	r.Run(":8080")
}
