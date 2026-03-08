package middleware

import (
	"net/http"
	"time"

	"github.com/Neel-max-cpu/go-rate-limiter/internal/redis"
	"github.com/gin-gonic/gin"
)

func RateLimit(limit int64, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		key := "rate_limit" + apiKey

		count, err := redis.Client.Incr(redis.Ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Redis error!",
			})
			c.Abort()
			return
		}

		if count == 1 {
			// When count is 1 (first request in the window), set the key to expire in 1 minute
			redis.Client.Expire(redis.Ctx, key, time.Minute)
		}

		if count > limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":  "rate limit exceeded!",
				"limit":  limit,
				"window": window.String(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
