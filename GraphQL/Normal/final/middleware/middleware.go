//This is the context for database
package middlewares

import (
	"context"
	"final/db"
	"net/http"
)

func DbContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//This will call function from file db and connect to the database
		db := database.Connect()
		ctx := context.WithValue(r.Context(), "database", db)
		defer db.Close()
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

	