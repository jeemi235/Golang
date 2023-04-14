package main

//! Manually set a key-value pair.

// import (
//   "github.com/bluele/gcache"
//   "fmt"
// )

// func main() {
//   gc := gcache.New(20).
//     LRU().
//     Build()
//   gc.Set("key", "ok")
//   value, err := gc.Get("key")
//   if err != nil {
//     panic(err)
//   }
//   fmt.Println("Get:", value)
// }

//! Manually set a key-value pair, with an expiration time.

// import (
// 	"github.com/bluele/gcache"
// 	"fmt"
// 	"time"
//   )

//   func main() {
// 	var err error
// 	gc := gcache.New(20).
// 	  LRU().
// 	  Build()
// 	gc.SetWithExpire("key", "ok", time.Second*10)
// 	value, _ := gc.Get("key")
// 	fmt.Println("Get:", value)

// 	// Wait for value to expire
// 	time.Sleep(time.Second*10)

// 	value, err = gc.Get("key")
// 	if err != nil {
// 	  panic(err)
// 	}
// 	fmt.Println("Get:", value)
//   }

//! Automatically load value

// import (
// 	"github.com/bluele/gcache"
// 	"fmt"
//   )

//   func main() {
// 	gc := gcache.New(20).
// 	  LRU().
// 	  LoaderFunc(func(key interface{}) (interface{}, error) {
// 		return "ok", nil
// 	  }).
// 	  Build()
// 	value, err := gc.Get("key")
// 	if err != nil {
// 	  panic(err)
// 	}
// 	fmt.Println("Get:", value)
//   }

//! Automatically load value with expiration

// import (
// 	"fmt"

// 	"time"

// 	"github.com/bluele/gcache"
//   )

//   func main() {
// 	var evictCounter, loaderCounter, purgeCounter int
// 	gc := gcache.New(20).
// 	  LRU().
// 	  LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
// 		loaderCounter++
// 		expire := 1 * time.Second
// 		return "ok", &expire, nil
// 	  }).
// 	  EvictedFunc(func(key, value interface{}) {
// 		evictCounter++
// 		fmt.Println("evicted key:", key)
// 	  }).
// 	  PurgeVisitorFunc(func(key, value interface{}) {
// 		purgeCounter++
// 		fmt.Println("purged key:", key)
// 	  }).
// 	  Build()
// 	value, err := gc.Get("key")
// 	if err != nil {
// 	  panic(err)
// 	}
// 	fmt.Println("Get:", value)
// 	time.Sleep(1 * time.Second)
// 	value, err = gc.Get("key")
// 	if err != nil {
// 	  panic(err)
// 	}
// 	fmt.Println("Get:", value)
// 	gc.Purge()
// 	if loaderCounter != evictCounter+purgeCounter {
// 	  panic("bad")
// 	}
//   }

//!LFU

// import "github.com/bluele/gcache"

// func main() {
// 	// size: 10
// 	gc := gcache.New(10).
// 		LFU().
// 		Build()
// 	gc.Set("key", "value")
// }

//!LRU

// import "github.com/bluele/gcache"

// func main() {
// 	// size: 10
// 	gc := gcache.New(10).
// 	  LRU().
// 	  Build()
// 	gc.Set("key", "value")
//   }

//!ARC

// import "github.com/bluele/gcache"

// func main() {
// 	// size: 10
// 	gc := gcache.New(10).
// 		ARC().
// 		Build()
// 	gc.Set("key", "value")
// }

//! SimpleCache: it has no clear priority for evict cache. It depends on key-value map order.

// import (
// 	"fmt"

// 	"github.com/bluele/gcache"
// )

// func main() {
// 	// size: 10
// 	gc := gcache.New(10).Build()
// 	gc.Set("key", "value")
// 	v, err := gc.Get("key")
// 	if err != nil {
// 	  panic(err)
// 	}
// 	fmt.Println("Get:",v)
//   }

//! Autoloading cache example

import (
	"fmt"
	"github.com/bluele/gcache"
)

func main() {
	gc := gcache.New(10).
		LFU().
		LoaderFunc(func(key interface{}) (interface{}, error) {
		return fmt.Sprintf("%v-value", key), nil
	}).
		Build()

	v, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}