package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var wg sync.WaitGroup

type Object struct {
	Str string
	Num int
}

func main() {
	defer fmt.Println("")
	wg.Add(2)
	go Example_advancedUsage()
	go Example_basicUsage()
	wg.Wait()
}

func Example_basicUsage() {
	defer wg.Done()
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6380",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		log.Println(err)
	}

	var wanted Object
	if err := mycache.Get(ctx, key, &wanted); err == nil {
		fmt.Println("\nBasic usage output:", wanted)
	}
}

func Example_advancedUsage() {
	defer wg.Done()
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6380",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	obj := new(Object)
	err := mycache.Once(&cache.Item{
		Key:   "mykey",
		Value: obj, // destination
		Do: func(*cache.Item) (interface{}, error) {
			return &Object{
				Str: "mystring",
				Num: 42,
			}, nil
		},
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("\nAdvance usage output:", obj)
}
