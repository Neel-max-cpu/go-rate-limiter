package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Connect() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	Client = redis.NewClient(&redis.Options{
		// Addr: "localhost:6379",
		Addr: host + ":" + port,
	})
}
