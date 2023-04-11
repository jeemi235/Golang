package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func main() {
	users := []User{
		{1, "jeemi", "455455"},
		{2, "sam", "484226"},
		{3, "chris", "782524"},
	}

	content, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	//write users data into json file
	err = ioutil.WriteFile("userfile.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//read data from json file
	content, err = ioutil.ReadFile("userfile.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	http.HandleFunc("/", handler)
	http.HandleFunc("/data", singledata)
	http.HandleFunc("/all", alldata)
	http.HandleFunc("/update", update)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello World")
}

// data of one user
func singledata(w http.ResponseWriter, r *http.Request) {

	user := User{Id: 1,
		Name:  "jeemi",
		Phone: "457454"}
	json.NewEncoder(w).Encode(user)

}

// data of all the users
func alldata(w http.ResponseWriter, r *http.Request) {

	users := []User{
		{1, "jeemi", "455455"},
		{2, "sam", "484226"},
		{3, "chris", "782524"},
	}
	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(out))
}

// updated data
func update(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{1, "Thakkar", "455455"},
		{2, "karan", "484226"},
		{3, "jordan", "782524"},
	}
	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	//json.NewEncoder(w).Encode(string(out))
	w.Write([]byte(out))
}
