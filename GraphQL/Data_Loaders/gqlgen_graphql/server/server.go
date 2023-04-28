package main

import (
	"gqlgen_graphql/dataloaders"
	"gqlgen_graphql/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const defaultPort = "5058"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//r.Use(graph.DbMiddleware)
	r.Use(dataloaders.ApplicationLoaderMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.Handle("/", playground.Handler("GraphQL plaground: ", "/query"))
	r.Handle("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})))
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
