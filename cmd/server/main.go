package main

// This defines an executable program in Go. The main package is the entry point for the application.

import (
	"log"
	"os"

	"github.com/Neel-max-cpu/go-rate-limiter/internal/db"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/redis"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/routes"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// This imports the Gin web framework, which provides all the tools for creating the web server and handling requests.

func main() {
	// load env
	// config.LoadEnv()

	// /*
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	// */

	// postgres connect
	err = db.Connect()
	// err := db.Connect()

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

	// all routes register
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default if not set
		log.Println("PORT not set, defaulting to", port)
	}
	// Run server on the configured port
	//log.Printf("Server starting on port %s", port)
	utils.InitLogger()
	utils.Logger.Info("Server starting",
		zap.String("port", port),
	)
	r.Run(":" + port)
}
