package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jakubnoga/pagila-graphql/graph/generated"
	"github.com/jakubnoga/pagila-graphql/graph/model"
)

// Actors is the resolver for the actors field.
func (r *filmResolver) Actors(ctx context.Context, obj *model.Film) ([]*model.Actor, error) {
	actors, err := r.ar.FindActorsByFilmId(obj.ID)

	if err != nil {
		panic(err)
	}

	result := make([]*model.Actor, len(actors))

	for idx, actor := range actors {
		result[idx] = actor.GraphModel()
	}

	return result, nil
}

// Categories is the resolver for the categories field.
func (r *filmResolver) Categories(ctx context.Context, obj *model.Film) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

// Film is the resolver for the film field.
func (r *queryResolver) Film(ctx context.Context, title string) (*model.Film, error) {
	film, err := r.fr.GetFilmsByTitle(title)

	return film.GraphModel(), err
}

// Films is the resolver for the films field.
func (r *queryResolver) Films(ctx context.Context) ([]*model.Film, error) {
	var result []*model.Film
	films, err := r.fr.GetFilms()

	if err != nil {
		return result, err
	}

	for _, film := range films {
		result = append(result, film.GraphModel())
	}

	return result, nil
}

// Film returns generated.FilmResolver implementation.
func (r *Resolver) Film() generated.FilmResolver { return &filmResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type filmResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
