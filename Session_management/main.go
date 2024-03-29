package main

import (
	"io"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

func main() {
	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	// Wrap your handlers with the LoadAndSave() middleware.
	http.ListenAndServe(":4000", sessionManager.LoadAndSave(mux))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	sessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := sessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}

/*
func loginHandler(w http.ResponseWriter, r *http.Request) {
	userID := 123

	// First renew the session token...
	err := sessionManager.RenewToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Then make the privilege-level change.
	sessionManager.Put(r.Context(), "userID", userID)
}


*/