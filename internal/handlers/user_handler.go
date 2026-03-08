package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Neel-max-cpu/go-rate-limiter/internal/db"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/dto"
	"github.com/Neel-max-cpu/go-rate-limiter/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest                  // 1. Create empty struct of expected shape
	if err := c.ShouldBindJSON(&req); err != nil { // 2. Fill it with request data
		// 3. Handle errors if request doesn't match expected shape
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	// creates id--
	id := uuid.New().String()
	// creates apikey--
	apiKey := uuid.New().String()

	query := `
	INSERT INTO users (id, email, api_key, created_at)
	VALUES ($1, $2, $3, $4)
	`
	// execute data
	_, err := db.DB.Exec(
		c,
		query,
		id,
		req.Email,
		apiKey,
		time.Now(),
	)

	if err != nil {
		// log
		log.Printf("DB insert failed: %v and error.Error():%v", err, err.Error())

		// Check if it's a duplicate email (more likely than duplicate UUID)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
			return
		}

		utils.Logger.Error("database error",
			zap.Error(err),
		)
		//return
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	// logger
	utils.Logger.Info("creating user",
		zap.String("email", req.Email),
	)

	c.JSON(http.StatusCreated, gin.H{
		"api_key": apiKey,
	})
}
