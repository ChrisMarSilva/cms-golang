package infra

// import (
//   "log"
//   "context"
//   "time"

//   "github.com/go-redis/redis/v8"
// )

// func NewRedisOrFatal(addr string) *redis.Client {
//   client := redis.NewClient(&redis.Options{Addr: addr})
//   if err := client.Ping(context.Background()).Err(); err != nil {
//     log.Fatal("Redis connect:", err)
//   }
//   return client
// }

// func SetHealthy(ctx context.Context, rdb *redis.Client, key string, healthy bool) {
//     val := "0"
//     if healthy {
//         val = "1"
//     }
//     rdb.Set(ctx, key, val, 5*time.Second)
// }

// func IsHealthy(ctx context.Context, rdb *redis.Client, key string) bool {
//     val, err := rdb.Get(ctx, key).Result()
//     return err == nil && val == "1"
// }
