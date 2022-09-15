package actors

import (
	"github.com/jakubnoga/pagila-graphql/actors/model"
	filmsModel "github.com/jakubnoga/pagila-graphql/films/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ActorsPostgresRepository struct {
	db *gorm.DB
}

func NewActorsPostgresRepository(dsn string) ActorsPostgresRepository {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return ActorsPostgresRepository{
		db,
	}
}

func (r *ActorsPostgresRepository) FindActorsByFilmId(filmId string) ([]model.Actor, error) {
	var film filmsModel.Film

	err := r.db.Model(&filmsModel.Film{}).Preload("Actors").First(&film, "film_id = ?", filmId).Error
	if err != nil {
		return []model.Actor{}, err
	}

	return film.Actors, nil
}
