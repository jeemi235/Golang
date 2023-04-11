package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	fs := http.FileServer(http.Dir("/Golang-Training/mux"))

// 	log.Fatal(http.ListenAndServe(":9090", fs))
// }


func main() {

	fs := http.FileServer(http.Dir("/Golang-Training/mux"))

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Golang!</h1>")
	})

	http.Handle( "/static/", http.StripPrefix( "/static", fs ) )

	log.Fatal(http.ListenAndServe(":9000", nil))

}
