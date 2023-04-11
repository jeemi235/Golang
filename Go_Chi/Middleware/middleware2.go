package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	router.Use(Logger) 

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello World"))
		if err != nil {
			log.Println(err)
		}
	})

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handler.ServeHTTP(writer, request)
	})
}