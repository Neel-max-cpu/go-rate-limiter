package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func Connect() {
	Client = redis.NewClient(&redis.Options{
		// Addr: "localhost:6379",
		Addr: "localhost:" + os.Getenv("REDIS_PORT"),
	})
}
