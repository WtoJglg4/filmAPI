package service

import "github/film-lib/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
