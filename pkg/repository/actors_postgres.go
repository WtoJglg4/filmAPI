package repository

import (
	"errors"
	"fmt"
	filmapi "github/film-lib"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ActorsPostgres struct {
	db *sqlx.DB
}

func NewActorsPostgres(db *sqlx.DB) *ActorsPostgres {
	return &ActorsPostgres{db: db}
}

func (r *ActorsPostgres) CreateActor(actor filmapi.Actor) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, to_date($3, 'DD.MM.YYYY')) RETURNING id", actorsTable)
	row := r.db.QueryRow(query, actor.Name, actor.Gender, actor.BirthDate)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ActorsPostgres) GetActorsList() ([]filmapi.ActorWithFilms, error) {
	var actors_list []filmapi.Actor
	query := fmt.Sprintf("SELECT id, name, gender, to_char(birth_date, 'DD.MM.YYYY') AS birth_date FROM %s", actorsTable)
	if err := r.db.Select(&actors_list, query); err != nil {
		return nil, err
	}
	actorsWithFilms := make([]filmapi.ActorWithFilms, len(actors_list))

	for i := range actors_list {
		actorsWithFilms[i].Actor = actors_list[i]
		query = "SELECT films.name FROM films JOIN actors_films ON actors_films.film_id = films.id JOIN actors ON actors_films.actor_id = actors.id WHERE actors.id = $1 ORDER BY films.rating DESC"
		if err := r.db.Select(&actorsWithFilms[i].Films, query, actors_list[i].Id); err != nil {
			return nil, err
		}
	}

	return actorsWithFilms, nil
}

func (r *ActorsPostgres) GetActorById(id int) (filmapi.ActorWithFilms, error) {
	var actorWithFilms []filmapi.ActorWithFilms
	query := fmt.Sprintf("SELECT id, name, gender, to_char(birth_date, 'DD.MM.YYYY') AS birth_date FROM %s WHERE id = %d", actorsTable, id)
	if err := r.db.Select(&actorWithFilms, query); err != nil {
		return filmapi.ActorWithFilms{}, err
	}

	if len(actorWithFilms) < 1 {
		return filmapi.ActorWithFilms{}, errors.New("sql: no actors with this id")
	}

	query = "SELECT films.name AS films_list FROM films JOIN actors_films ON actors_films.film_id = films.id JOIN actors ON actors_films.actor_id = actors.id WHERE actors.id = $1 ORDER BY films.rating DESC"
	if err := r.db.Select(&actorWithFilms[0].Films, query, actorWithFilms[0].Id); err != nil {
		return filmapi.ActorWithFilms{}, err
	}

	return actorWithFilms[0], nil
}

func (r *ActorsPostgres) UpdateActorById(name, gender, birthDate string, id int) error {
	args := make([]string, 0)
	if name != "" {
		args = append(args, fmt.Sprintf("name = '%s' ", name))
	}
	if gender != "" {
		args = append(args, fmt.Sprintf("gender = '%s' ", gender))
	}
	if birthDate != "" {
		args = append(args, fmt.Sprintf("birth_date = to_date('%s', 'DD.MM.YYYY')", birthDate))
	}
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", actorsTable, strings.Join(args, ","), id)

	_, err := r.db.Exec(query)
	return err
}

func (r *ActorsPostgres) DeleteActorById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=%d", actorsTable, id)
	_, err := r.db.Exec(query)
	return err
}
