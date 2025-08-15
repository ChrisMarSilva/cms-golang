package stores

import (
	"context"
	"log"
	"time"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
	//ttl    time.Duration
}

func NewRedisCache(config *utils.Config) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
		//Username:     config.RedisUser,
		Password:        config.RedisPwd,
		DB:              0,   // use default DB
		PoolSize:        100, // max-active
		MaxIdleConns:    200,
		MinIdleConns:    100,
		ConnMaxIdleTime: time.Duration(300) * time.Second,
		ReadTimeout:     time.Duration(5) * time.Second,
		WriteTimeout:    time.Duration(5) * time.Second,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Redis connect:", err)
	}

	//return client
	return &RedisCache{
		Client: client,
		//ttl:    5 * time.Minute,
	}
}

func (r *RedisCache) HSet(ctx context.Context, key, id string, data interface{}) error {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HSet")
	defer span.End()

	return r.Client.HSet(ctx, key, id, data).Err()
}

func (r *RedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HGetAll")
	defer span.End()

	return r.Client.HGetAll(ctx, key).Result()
}

func (r *RedisCache) HGet(ctx context.Context, key, id string) (string, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HGet")
	defer span.End()

	return r.Client.HGet(ctx, key, id).Result()
}

func (r *RedisCache) Exists(ctx context.Context, key, id string) (bool, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.Exists")
	defer span.End()

	return r.Client.HExists(ctx, key, id).Result()
}

func (r *RedisCache) HLen(ctx context.Context, key string) (int64, error) {
	_, span := utils.Tracer.Start(ctx, "RedisCache.HLen")
	defer span.End()

	return r.Client.HLen(ctx, key).Result()
}

func (r *RedisCache) Delete(ctx context.Context, key, id string) error {
	_, span := utils.Tracer.Start(ctx, "RedisCache.Delete")
	defer span.End()

	return r.Client.HDel(ctx, key, id).Err()
}

func (r *RedisCache) Close() {
	r.Client.Close()
}
