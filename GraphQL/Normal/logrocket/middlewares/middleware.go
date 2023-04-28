package middlewares

import (
	"context"
	"logrocket/db"
	"net/http"
)

func DbContext(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.Connect()
		ctx := context.WithValue(r.Context(), "database", db)
		defer db.Close()
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
