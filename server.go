package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jakubnoga/pagila-graphql/actors"
	"github.com/jakubnoga/pagila-graphql/films"
	"github.com/jakubnoga/pagila-graphql/graph"
	"github.com/jakubnoga/pagila-graphql/graph/generated"
)

const defaultPort = "8080"
const dsn = "host=localhost user=postgres password=123456 dbname=postgres port=5432"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	filmsRepository := films.NewFilmsPostgresRepository(dsn)
	actorsRepository := actors.NewActorsPostgresRepository(dsn)
	resolver := graph.NewResolver(filmsRepository, actorsRepository)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
