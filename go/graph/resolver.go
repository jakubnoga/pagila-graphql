package graph

import (
	"github.com/jakubnoga/pagila-graphql/actors"
	"github.com/jakubnoga/pagila-graphql/films"
)

//go:generate go run github.com/99designs/gqlgen generate
// import "github.com/jakubnoga/pagila-graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	fr films.FilmsPostgresRepository
	ar actors.ActorsPostgresRepository
}

func NewResolver(filmsRepository films.FilmsPostgresRepository, actorsRepository actors.ActorsPostgresRepository) Resolver {
	return Resolver{filmsRepository, actorsRepository}
}
