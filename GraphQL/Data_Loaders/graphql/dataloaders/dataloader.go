package dataloader

import (
	"context"
	database "graphql/db"
	"graphql/graph/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type loaders struct {
	Todos *TodoLoader
}

func LoadMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		ctx := r.Context()
		// crConn := ctx.Value("crConn").(*sql.DB)
		crConn := database.Connect()
		wait := 350 * time.Millisecond

		// Todo loader
		ldrs.Todos = &TodoLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(ids []string) ([]*model.Todo, []error) {
				var sqlQuery string
				if len(ids) == 1 {
					sqlQuery = "SELECT id, name FROM todo WHERE id = ?"
				} else {
					sqlQuery = "SELECT id, name FROM todo WHERE id IN (?)"
				}
				sqlQuery, arguments, err := sqlx.In(sqlQuery, ids)
				if err != nil {
					log.Println("Error to fetch user query")
				}
				sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
				rows, err := crConn.Query(sqlQuery, arguments...)
				if err != nil {
					log.Println("Error to read user data")
				}
				todoById := map[int]*model.Todo{}
				defer func() {
					if err := rows.Close(); err != nil {
						log.Println("Error at defer function of user")
					}
				}()

				for rows.Next() {
					var todo model.Todo
					err := rows.Scan(&todo.ID, &todo.Name)
					if err != nil {
						log.Println("Error to scan todo data")
					}
					// Store user.ID in key & model of user in value
					x, _ := strconv.Atoi(todo.ID)
					todoById[x]=&todo
				}

				todos := make([]*model.Todo, len(ids))
				for i, id := range ids {
					x, _ := strconv.Atoi(id)
					todos[i] = todoById[x]
					i++
				}
				return todos, nil
			},
		}

		ctx = context.WithValue(ctx, ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
