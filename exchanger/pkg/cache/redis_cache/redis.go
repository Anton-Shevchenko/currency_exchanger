package redis_cache

import (
	"context"
	"fmt"
	redisClient "github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type Cache struct {
	Client *redisClient.Client
}

func NewRedisCache() *Cache {
	return &Cache{
		Client: redisClient.NewClient(&redisClient.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (redis Cache) Set(key string, data any, ttl time.Duration) error {
	err := redis.Client.Set(ctx, key, data, ttl).Err()

	if err != nil {
		fmt.Println(err)

		return err
	}

	return nil
}

func (redis Cache) Get(key string) (value string, isExists bool) {
	val, err := redis.Client.Get(ctx, key).Result()

	if err == redisClient.Nil {
		return "", false
	} else if err != nil {
		panic(err)
	} else {
		return val, true
	}
}

func (redis Cache) Exists(key string) int64 {
	val, err := redis.Client.Exists(ctx, key).Result()

	if err != nil {
		panic(err)
	} else {
		return val
	}
}

func (redis Cache) Keys(key string) []string {
	val, err := redis.Client.Keys(ctx, key).Result()

	if err != nil {
		panic(err)
	} else {
		return val
	}
}

func (redis Cache) Delete(key string) error {
	_, err := redis.Client.Del(ctx, key).Result()

	return err
}
