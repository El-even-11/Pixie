package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisDB *redis.Client

func RedisInit() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Set(key string, value interface{}) error {
	return redisDB.Set(ctx, key, value, 0).Err()
}

func Get(key string) (string, error) {
	return redisDB.Get(ctx, key).Result()
}

func SAdd(key string, members []string) error {
	t := make([]interface{}, len(members))
	for i, v := range members {
		t[i] = v
	}
	return redisDB.SAdd(ctx, key, t...).Err()
}

func SMembers(key string) ([]string, error) {
	return redisDB.SMembers(ctx, key).Result()
}
