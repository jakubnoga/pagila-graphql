package films

import (
	"github.com/jakubnoga/pagila-graphql/films/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=123456 dbname=postgres port=5432"

type FilmsPostgresRepository struct {
	db *gorm.DB
}

func NewFilmsPostgresRepository(dsn string) FilmsPostgresRepository {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return FilmsPostgresRepository{
		db,
	}
}

func (r *FilmsPostgresRepository) GetFilmsByTitle(title string) (*model.Film, error) {
	var film model.Film

	err := r.db.First(&film, "title = ?", title).Error

	return &film, err
}

func (r *FilmsPostgresRepository) GetFilms() ([]*model.Film, error) {
	var films []*model.Film

	err := r.db.Find(&films).Error

	return films, err
}
