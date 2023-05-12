package dataloaders

import (
	"context"
	"fmt"
	"gqlgen_graphql/graph/model"
	"log"
	"net/http"
	"time"

	database "gqlgen_graphql/db"

	"github.com/jmoiron/sqlx"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type loaders struct {
	Todos *TodoLoader
}

// ApplicationLoaderMiddleware for datloader
func ApplicationLoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ldrs := loaders{}
		ctx := request.Context()
		// crConn := ctx.Value("crConn").(*sql.DB)
		crConn := database.Connect()
		wait := 350 * time.Millisecond

		// Todo loader
		ldrs.Todos = &TodoLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(ids []string) ([][]*model.Todo, []error) {
				var sqlQuery string
				fmt.Println(len(ids))
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
				fmt.Println(sqlQuery)
				var todoById []*model.Todo
				//todoById := map[int]*model.Todo{}
				defer func() {
					if err := rows.Close(); err != nil {
						log.Println("Error at defer function of user")
					}
				}()

				for rows.Next() {
					var todo model.Todo
					err := rows.Scan(&todo.ID, &todo.Name)
					if err != nil {
						log.Println("Error to scan user data")
					}
					// Store todo.ID in key & model of user in value
					todoById = append(todoById, &todo)
				}
				fmt.Println(todoById)

				todos := make([]*model.Todo, len(ids))
				// for i, id := range ids {
				// 	x, _ := strconv.Atoi(id)
				// 	todos[i] = todoById[x]
				// 	i++
				// }
				todos = append(todos, todoById...)
				fmt.Println("todos:")
				fmt.Println(todos)
				return [][]*model.Todo{}, nil
			},
		}
		ctx = context.WithValue(ctx, ctxKey, ldrs)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
