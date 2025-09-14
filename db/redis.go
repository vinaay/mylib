package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisClient{Client: rdb}
}

func (r *RedisClient) SetValue(ctx context.Context, key, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *RedisClient) GetValue(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
