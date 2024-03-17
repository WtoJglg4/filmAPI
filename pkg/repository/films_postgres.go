package repository

import (
	"errors"
	"fmt"
	filmapi "github/film-lib"
	"strings"

	"github.com/jmoiron/sqlx"
)

type FilmsPostgres struct {
	db *sqlx.DB
}

func NewFilmsPostgres(db *sqlx.DB) *FilmsPostgres {
	return &FilmsPostgres{db: db}
}

func (r *FilmsPostgres) CreateFilm(film filmapi.Film) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, description, release_date, rating) VALUES ($1, $2, to_date($3, 'DD.MM.YYYY'), $4) RETURNING id", filmsTable)
	row := tx.QueryRow(query, film.Name, film.Description, film.ReleaseDate, film.Rating)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, actor_name := range film.Actors {
		var actor_id int
		query = fmt.Sprintf("SELECT id FROM %s WHERE name=$1", actorsTable)
		row := tx.QueryRow(query, actor_name)
		if err := row.Scan(&actor_id); err != nil {
			tx.Rollback()
			return 0, errors.New(err.Error() + ": unknown actor")
		}

		query = fmt.Sprintf("INSERT INTO %s (actor_id, film_id) VALUES (%d, %d)", actorsFilmsTable, actor_id, id)
		_, err := tx.Exec(query)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return id, tx.Commit()
}

func (r *FilmsPostgres) GetFilmsList(sort string) ([]filmapi.Film, error) {
	var order string = "DESC"
	if sort == "name" {
		order = "ASC"
	}
	fmt.Println(sort, order)

	var films_list []filmapi.Film

	query := fmt.Sprintf("SELECT id, name, description, to_char(release_date, 'DD.MM.YYYY') AS release_date_new, rating FROM %s ORDER BY %s %s", filmsTable, sort, order)
	if err := r.db.Select(&films_list, query); err != nil {
		return nil, err
	}
	for i, f := range films_list {
		query = "SELECT actors.name AS actors_list FROM actors JOIN actors_films ON actors_films.actor_id = actors.id JOIN films ON actors_films.film_id = films.id WHERE films.id = $1"
		if err := r.db.Select(&films_list[i].Actors, query, f.Id); err != nil {
			return nil, err
		}
	}
	return films_list, nil
}

func (r *FilmsPostgres) UpdateFilmById(name, description, releaseDate string, rating, id int) error {
	args := make([]string, 0)
	if name != "" {
		args = append(args, fmt.Sprintf("name = '%s' ", name))
	}
	if description != "" {
		args = append(args, fmt.Sprintf("description = '%s' ", description))
	}
	if releaseDate != "" {
		args = append(args, fmt.Sprintf("release_date = to_date('%s', 'DD.MM.YYYY') ", releaseDate))
	}
	if rating != 0 {
		args = append(args, fmt.Sprintf("rating = %d", rating))
	}
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", filmsTable, strings.Join(args, ","), id)

	_, err := r.db.Exec(query)
	return err
}

func (r *FilmsPostgres) DeleteFilmById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=%d", filmsTable, id)
	_, err := r.db.Exec(query)
	return err
}

func (r *FilmsPostgres) GetFilmByPart(parameter, req string) ([]filmapi.Film, error) {
	var films_list []filmapi.Film
	var query string

	if parameter == "actor" {
		subquery := fmt.Sprintf("SELECT COUNT(actors.id) FROM actors JOIN actors_films af ON af.actor_id = actors.id JOIN films f ON f.id = af.film_id WHERE f.id = films.id AND actors.name LIKE '%%%s%%'", req)
		query = fmt.Sprintf("SELECT id, name, description, to_char(release_date, 'DD.MM.YYYY') AS release_date_new, rating FROM %s WHERE(%s) > 0", filmsTable, subquery)
	} else {
		query = fmt.Sprintf("SELECT id, name, description, to_char(release_date, 'DD.MM.YYYY') AS release_date_new, rating FROM %s WHERE name LIKE '%%%s%%'", filmsTable, req)
	}
	if err := r.db.Select(&films_list, query); err != nil {
		return nil, err
	}
	for i, f := range films_list {
		query = "SELECT actors.name AS actors_list FROM actors JOIN actors_films ON actors_films.actor_id = actors.id JOIN films ON actors_films.film_id = films.id WHERE films.id = $1"
		if err := r.db.Select(&films_list[i].Actors, query, f.Id); err != nil {
			return nil, err
		}
	}
	return films_list, nil
}
