package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println("booting on localhost:6050")
	log.Fatal(http.ListenAndServe(":6050", nil))
}
