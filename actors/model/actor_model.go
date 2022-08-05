package model

import graphModel "github.com/jakubnoga/pagila-graphql/graph/model"

type Actor struct {
	ID        string `gorm:"column:actor_id"`
	FirstName string
	LastName  string
}

func (Actor) TableName() string {
	return "actor"
}

func (a *Actor) GraphModel() *graphModel.Actor {
	return &graphModel.Actor{
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}
