package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)

	go func(ch chan string) {
		for {
			Cache.Del(context.Background(), <-ch)

			fmt.Println("Cache Cleared")
		}
	}(CacheChannel)
}

func ClearCache(key string) {
	CacheChannel <- key
}
