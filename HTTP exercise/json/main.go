package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string
	Password string
}

func main() {
	users := []user{
		{"debora", "123456"},
		{"bob", "42"},
		{"sandra", "33"},
	}

	// use marshal func to convert to json
	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print encoded json data
	fmt.Println(string(out))
}
