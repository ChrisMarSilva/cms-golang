package utils

import (
	// _ "fmt"
	// _ "os"

	"github.com/go-redis/redis/v7"
)

//NewRedisClient create a new instace of client redis
func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
		Password: "123",
		DB:       4, // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

}
