package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Todos struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var err error
	db, err = sql.Open("postgres", "postgresql://max:roach@localhost:26257/task?sslmode=require")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/todo", posttodo)
	r.Get("/todo", gettodos)

	fmt.Println("server started at http port :5051")
	http.ListenAndServe(":5051", r)
}

func posttodo(w http.ResponseWriter, r *http.Request) {
	var Todo Todos
	err := json.NewDecoder(r.Body).Decode(&Todo)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	if _, err := db.Exec(
		"INSERT INTO todo(id,name) values($1,$2)", Todo.Id, Todo.Name); err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Data inserted succesfully successfully"))
	log.Print("successfully inserted todo details\n")
}

func gettodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from todo")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	Todo := []Todos{}
	for rows.Next() {
		var tod Todos
		if err := rows.Scan(&tod.Id, &tod.Name); err != nil {
			log.Println(err)
		}
		Todo = append(Todo, tod)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(&Todo)
	fmt.Print("successfully got todo details\n")
}
