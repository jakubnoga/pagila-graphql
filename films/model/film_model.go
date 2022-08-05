package model

import actorsModel "github.com/jakubnoga/pagila-graphql/actors/model"
import graphModel "github.com/jakubnoga/pagila-graphql/graph/model"

type Film struct {
	ID          string `gorm:"column:film_id"`
	Title       string
	Description string
	Actors      []actorsModel.Actor `gorm:"many2many:film_actor"`
}

func (Film) TableName() string {
	return "film"
}

func (f *Film) GraphModel() *graphModel.Film {
	var actors []*graphModel.Actor

	for _, actor := range f.Actors {
		actors = append(actors, actor.GraphModel())
	}

	return &graphModel.Film{
		ID:          f.ID,
		Title:       f.Title,
		Description: f.Description,
		Actors:      actors,
	}
}
