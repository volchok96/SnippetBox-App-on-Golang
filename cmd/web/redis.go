package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func connectToRedis(addr string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}