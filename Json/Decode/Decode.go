package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var human1 Human

	// data in JSON format which is to be decoded
	Data := []byte(`{
		"Name": "Deeksha",
		"Address": "Hyderabad",
		"Age": 21
	}`)

	// decoding human1 struct from json format
	err := json.Unmarshal(Data, &human1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Struct is:", human1)
	fmt.Printf("%s lives in %s.\n", human1.Name, human1.Address)

	// unmarshalling a JSON array to array type in Golang

	// defining an array instance of struct type
	var human2 []Human

	// JSON array to be decoded to an array
	Data2 := []byte(`
	[
		{"Name": "Vani", "Address": "Delhi", "Age": 21},
		{"Name": "Rashi", "Address": "Noida", "Age": 24},
		{"Name": "Rohit", "Address": "Pune", "Age": 25}
	]`)

	// decoding JSON array to human2 array
	err2 := json.Unmarshal(Data2, &human2)

	if err2 != nil {
		fmt.Println(err2)
	}
	// printing decoded array values one by one
	for i := range human2 {
		fmt.Println(human2[i])
	}
}

// while unstructure data

// str := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

// var reading map[string]interface{}
// err = json.Unmarshal([]byte(str), &reading)
// fmt.Printf("%+v\n", reading)
