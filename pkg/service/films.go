package service

import (
	filmapi "github/film-lib"
	"github/film-lib/pkg/repository"
)

type FilmService struct {
	repo repository.Film
}

func NewFilmService(repo repository.Film) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) CreateFilm(film filmapi.Film) (int, error) {
	return s.repo.CreateFilm(film)
}

func (s *FilmService) GetFilmsList(sort string) ([]filmapi.Film, error) {
	return s.repo.GetFilmsList(sort)
}

func (s *FilmService) UpdateFilmById(name, description, releaseDate string, rating, id int) error {
	return s.repo.UpdateFilmById(name, description, releaseDate, rating, id)
}

func (s *FilmService) DeleteFilmById(id int) error {
	return s.repo.DeleteFilmById(id)
}

func (s *FilmService) GetFilmByPart(parameter, req string) ([]filmapi.Film, error) {
	return s.repo.GetFilmByPart(parameter, req)
}
