package main

import (
	"time"

	"log"
	"net/http"

	"sat/handler"
	"sat/store"

	"github.com/bluele/gcache"
	"github.com/gorilla/mux"
)

func main() {

	db := map[string]int{
		"Tom":   1410,
		"Jill":  1420,
		"Mike":  1290,
		"Alice": 1450,
	}

	store := store.NewStore(db, gcache.New(30).LFU().Build(), time.Hour)
	sh := handler.NewScoreHandler(store)
	r := mux.NewRouter()
	r.HandleFunc("/scores/{student}", sh.HandleGet).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// http://127.0.0.1:8000/scores/Tom
