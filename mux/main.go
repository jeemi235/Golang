package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {

    r := mux.NewRouter()

    r.HandleFunc("/hello", func(resp http.ResponseWriter, _ *http.Request) {

        fmt.Fprint(resp, "Hello there!")
    })

    log.Println("Listening...")
    http.ListenAndServe(":5555", r)
}

// func main() {

//     r := mux.NewRouter()

//     r.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {

//         name := req.URL.Query().Get("name")

//         if name == "" {
//             name = "guest"
//         }

//         fmt.Fprintf(resp, "Hello %s!", name)
//     })

//     log.Println("Listening ...")
//     http.ListenAndServe(":5555", r)
// }

// func main() {

//     r := mux.NewRouter()

//     r.HandleFunc("/hello/{name}", func(resp http.ResponseWriter, req *http.Request) {

//         vars := mux.Vars(req)
//         name := vars["name"]

//         fmt.Fprintf(resp, "Hello %s!", name)
//     })

//     log.Println("Listening ...")
//     http.ListenAndServe(":5555", r)
// }

// func main() {

// 	r := mux.NewRouter()
// 	r.HandleFunc("/now", NowHandler)

// 	log.Println("Listening ...")
// 	http.ListenAndServe(":5555", r)
// }

// func NowHandler(resp http.ResponseWriter, _ *http.Request) {

// 	now := time.Now()

// 	payload := make(map[string]string)
// 	payload["now"] = now.Format(time.ANSIC)

// 	resp.Header().Set("Content-Type", "application/json")
// 	resp.WriteHeader(http.StatusOK)

// 	json.NewEncoder(resp).Encode(payload)
// }

// func main() {

// 	r := mux.NewRouter()

// 	s1 := r.PathPrefix("/path1").Subrouter()
// 	s1.HandleFunc("/", Handler1)

// 	s2 := r.PathPrefix("/path2").Subrouter()
// 	s2.HandleFunc("/", Handler2)

// 	log.Println("Listening ...")
// 	http.ListenAndServe(":5555", r)
// }

// func Handler1(resp http.ResponseWriter, _ *http.Request) {

// 	fmt.Fprint(resp, "Subroute 1")
// }

// func Handler2(resp http.ResponseWriter, _ *http.Request) {

// 	fmt.Fprint(resp, "Subroute 2")
// }