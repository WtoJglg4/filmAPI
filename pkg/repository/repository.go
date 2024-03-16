package repository

import (
	filmapi "github/film-lib"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user filmapi.User) (int, error)
	GetUser(username, password string) (filmapi.User, error)
}

type Actor interface {
	CreateActor(actor filmapi.Actor) (int, error)
	GetActorsList() ([]filmapi.ActorWithFilms, error)
	GetActorById(id int) (filmapi.ActorWithFilms, error)
	UpdateActorById(name, gender, birthDate string, id int) error
	DeleteActorById(id int) error
}

type Film interface {
}

type Repository struct {
	Authorization
	Actor
	Film
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Actor:         NewActorsPostgres(db),
	}
}
