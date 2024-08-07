package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func connectToRedis(addr string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr, // Address of the Redis server
	})

	// Check the connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
