package stores

import (
	"context"
	"log"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/internal/utils"
	"github.com/redis/go-redis/v9"
)

func NewRedis(config *utils.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         config.RedisAddr,
		Password:     config.RedisPwd,
		DB:           0,
		PoolSize:     100,
		MinIdleConns: 25,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Redis connect:", err)
	}

	return client
}

// func SetHealthy(ctx context.Context, rdb *redis.Client, key string, healthy bool) {
// 	rdb.Set(ctx, key, val, 5*time.Second)
// }

// func IsHealthy(ctx context.Context, rdb *redis.Client, key string) bool {
// 	val, err := rdb.Get(ctx, key).Result()
// 	return err == nil && val == "1"
// }
