package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello World"))
	if err != nil {
		log.Println(err)
	}
}

// path params using chi.urlparam method
func getusername(writer http.ResponseWriter, request *http.Request) {
	username := chi.URLParam(request, "username")
	_, err := writer.Write([]byte("Hello " + username))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", hello)
	router.Get("/user/{username}", getusername)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}
