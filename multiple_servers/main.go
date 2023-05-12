package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {

	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	// create a default route handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello world: "+req.Host)
	})

	// goroutine to launch a server on port 9000
	go func() {
		fmt.Println("running server on: 9000")
		log.Fatal(http.ListenAndServe(":9000", nil))
		wg.Done() // one goroutine finished
	}()

	// goroutine to launch a server on port 9001
	go func() {
		fmt.Println("running server on: 9001")
		log.Fatal(http.ListenAndServe(":9001", nil))
		wg.Done() // one goroutine finished
	}()

	// wait until WaitGroup is done
	wg.Wait()

}
