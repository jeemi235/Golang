package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", HandleAddInts)
	address := ":9000"
	fmt.Printf("Starting math server at %s\n", address)
	fmt.Println()
	fmt.Println("Example query: /add?a=2&b=2&authtoken=abcdef123")
	fmt.Println("The auth token is 'abcdef123'")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("error while starting server: %v", err)
	}
}

type AddResponse struct {
	Result int `json:"result"`
}

func HandleAddInts(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	paramA := params.Get("a")
	paramB := params.Get("b")
	token := params.Get("authtoken")

	if paramA == "" || paramB == "" || token == "" {
		http.Error(w, "not enough parameters", http.StatusBadRequest)
		return
	}

	if token != "abcdef123" {
		http.Error(w, "invalid auth token", http.StatusUnauthorized)
		return
	}

	intA, err := strconv.Atoi(paramA)
	if err != nil {
		http.Error(w, "parameter 'a' must be a valid integer", http.StatusBadRequest)
		return
	}

	intB, err := strconv.Atoi(paramB)
	if err != nil {
		http.Error(w, "parameter 'b' must be a valid integer", http.StatusBadRequest)
		return
	}

	response := AddResponse{
		Result: intA + intB,
	}

	json, err := json.MarshalIndent(&response, "", " ")
	if err != nil {
		http.Error(w, "error while marshalling", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(json))
}
