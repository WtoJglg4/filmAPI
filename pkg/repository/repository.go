package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
