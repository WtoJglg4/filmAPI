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
	}
}
