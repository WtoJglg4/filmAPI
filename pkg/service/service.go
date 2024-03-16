package service

import (
	filmapi "github/film-lib"
	"github/film-lib/pkg/repository"
)

type Authorization interface {
	CreateUser(user filmapi.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Actor interface {
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
	}
}
