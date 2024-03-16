package service

import (
	filmapi "github/film-lib"
	"github/film-lib/pkg/repository"
)

type ActorService struct {
	repo repository.Actor
}

func NewActorService(repo repository.Actor) *ActorService {
	return &ActorService{repo: repo}
}

func (s *ActorService) CreateActor(actor filmapi.Actor) (int, error) {
	return s.repo.CreateActor(actor)
}

func (s *ActorService) GetActorsList() ([]filmapi.ActorWithFilms, error) {
	return s.repo.GetActorsList()
}

func (s *ActorService) GetActorById(id int) (filmapi.ActorWithFilms, error) {
	return s.repo.GetActorById(id)
}

func (s *ActorService) UpdateActorById(name, gender, birthDate string, id int) error {
	return s.repo.UpdateActorById(name, gender, birthDate, id)
}

func (s *ActorService) DeleteActorById(id int) error {
	return s.repo.DeleteActorById(id)
}
