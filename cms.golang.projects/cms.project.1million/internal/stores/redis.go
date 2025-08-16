package stores

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	logger *slog.Logger
	client *redis.Client
}

func NewRedisCache(logger *slog.Logger, config *utils.Config) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:            config.RedisAddr,
		Password:        config.RedisPwd,
		DB:              0,
		PoolSize:        100,
		MaxIdleConns:    200,
		MinIdleConns:    100,
		ConnMaxIdleTime: time.Duration(300) * time.Second,
		ReadTimeout:     time.Duration(5) * time.Second,
		WriteTimeout:    time.Duration(5) * time.Second,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Error("Redis connection error", "error", err)
		os.Exit(1)
	}

	return &RedisCache{
		logger: logger,
		client: client,
	}
}

func (r *RedisCache) Set(ctx context.Context, key string, value string) error {
	_, span := utils.Tracer.Start(ctx, "RedisCache.Set")
	defer span.End()

	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.Get")
	defer span.End()

	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Exists(ctx context.Context, key string) (int64, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.Exists")
	defer span.End()

	return r.client.Exists(ctx, key).Result()
}

func (r *RedisCache) Del(ctx context.Context, key string) error {
	_, span := utils.Tracer.Start(ctx, "RedisCache.Del")
	defer span.End()

	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) HSet(ctx context.Context, key, id string, data interface{}) error {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HSet")
	defer span.End()

	return r.client.HSet(ctx, key, id, data).Err()
}

func (r *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HGetAll")
	defer span.End()

	return r.client.HGetAll(ctx, key).Result()
}

func (r *RedisCache) HGet(ctx context.Context, key, id string) (string, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HGet")
	defer span.End()

	return r.client.HGet(ctx, key, id).Result()
}

func (r *RedisCache) HExists(ctx context.Context, key, id string) (bool, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HExists")
	defer span.End()

	return r.client.HExists(ctx, key, id).Result()
}

func (r *RedisCache) HLen(ctx context.Context, key string) (int64, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HLen")
	defer span.End()

	return r.client.HLen(ctx, key).Result()
}

func (r *RedisCache) HDel(ctx context.Context, key, id string) error {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HDel")
	defer span.End()

	return r.client.HDel(ctx, key, id).Err()
}

func (r *RedisCache) Pipeline() *redis.Pipeline {
	return r.client.Pipeline().(*redis.Pipeline)
}

func (r *RedisCache) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
