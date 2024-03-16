package repository

import (
	"fmt"
	filmapi "github/film-lib"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable       = "users"
	actorsTable      = "actors"
	filmsTable       = "films"
	actorsFilmsTable = "actors_films"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user filmapi.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, role) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (filmapi.User, error) {
	var user filmapi.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
