package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Println("helo world")
	var wg sync.WaitGroup
	wg.Add(1)
	//go leakyFunction(wg)
	wg.Wait()
}

//curl http://localhost:6060/debug/pprof/heap

// http://localhost:6060/debug/pprof/heap?seconds=n

// curl http://localhost:6060/debug/pprof/heap --output heap.tar.gz

// go tool pprof heap.tar.gz

