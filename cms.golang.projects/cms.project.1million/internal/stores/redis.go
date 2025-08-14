package stores

import (
	"context"
	"log"
	"time"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/redis/go-redis/v9"
)

// type RedisCache struct {
// 	client *redis.Client
// 	ttl    time.Duration
// }

func NewRedis(config *utils.Config) *redis.Client {
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

	// return &RedisCache{
	// 	client: client,
	// 	ttl:    5 * time.Minute,
	// }
	return client
}

// func (r *RedisCache) Set(ctx context.Context, key, id string, data *interface{}) {
// 	key := fmt.Sprintf("pessoa:id:%d", p.ID)
// 	payload, _ := json.Marshal(data)
// 	return r.client.Set(ctx, key, payload, r.ttl)
//  return r.client.HSet(ctx, key, id, payload).Err()
// }

// func (r *RedisCache) GetAll(ctx context.Context, key string) (*interface{}, error) {
// 	return r.client.Get(ctx, key).Result()
//  return r.client.HGetAll(ctx, key).Result()
// }

// func (r *RedisCache) GetPessoaByID(ctx context.Context, id uuid) (*interface{}, error) {
// 	key := fmt.Sprintf("pessoa:id:%d", id.String())
// 	payload, err := r.client.Get(ctx, key).Result()
//  payload, err := r.client.HGet(ctx, key, id.String()).Result()
// 	if err == redis.Nil {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	var data *interface{
// 	if err := json.Unmarshal([]byte(payload), &data); err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }
