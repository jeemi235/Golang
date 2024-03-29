package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zenyui/gqlgen-dataloader/graph/dataloader"
	"github.com/zenyui/gqlgen-dataloader/graph/generated"
	"github.com/zenyui/gqlgen-dataloader/graph/resolver"
	"github.com/zenyui/gqlgen-dataloader/graph/storage"
)

const defaultPort = "5057"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// instantiate the DB client
	db := storage.NewMemoryStorage()
	// make a data loader
	loader := dataloader.NewDataLoader(db)
	// instantiate the gqlgen Graph Resolver
	graphResolver := resolver.NewResolver(db)
	// create the query handler
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graphResolver}))
	// wrap the query handler with middleware to inject dataloader
	dataloaderSrv := dataloader.Middleware(loader, srv)
	// register the query endpoint
	http.Handle("/query", dataloaderSrv)
	// register the playground
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// boot the server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
