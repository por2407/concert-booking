package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})

	// ลอง Ping ดูว่าเจอไหม
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to Redis successfully!")
	return rdb, nil
}
