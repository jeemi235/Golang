package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func main() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "mykey"
	obj := "1"

	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		log.Println(err)
	}

	if err := mycache.Get(ctx, key, &obj); err == nil {
		fmt.Println("output:", 1)
	}
}
