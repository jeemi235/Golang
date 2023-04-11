package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// jsonHandler returns http respone in JSON format.
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{Id: 1,
		Name:  "John Doe",
		Phone: "000099999"}
	json.NewEncoder(w).Encode(user)
}
