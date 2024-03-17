package service

import (
	filmapi "github/film-lib"
	"github/film-lib/pkg/repository"
)

type Authorization interface {
	CreateUser(user filmapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Actor interface {
	CreateActor(actor filmapi.Actor) (int, error)
	GetActorsList() ([]filmapi.ActorWithFilms, error)
	GetActorById(id int) (filmapi.ActorWithFilms, error)
	UpdateActorById(name, gender, birthDate string, id int) error
	DeleteActorById(id int) error
}

type Film interface {
	CreateFilm(film filmapi.Film) (int, error)
	GetFilmsList(sort string) ([]filmapi.Film, error)
	UpdateFilmById(name, description, releaseDate string, rating, id int) error
	DeleteFilmById(id int) error
	GetFilmByPart(parameter, req string) ([]filmapi.Film, error)
}

type Service struct {
	Authorization
	Actor
	Film
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Actor:         NewActorService(repo.Actor),
		Film:          NewFilmService(repo.Film),
	}
}
