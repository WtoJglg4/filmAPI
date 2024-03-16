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
	}
}
