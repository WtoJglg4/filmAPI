package filmapi

type Actor struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" bindig:"required" db:"name"`
	Gender    string `json:"gender" bindig:"required" db:"gender"`
	BirthDate string `json:"birth_date" bindig:"required" db:"birth_date"`
}

type Film struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" bindig:"required" db:"name"`
	Description string `json:"description" bindig:"required" db:"description"`
	ReleaseDate string `json:"release_date" bindig:"required" db:"release_date"`
	Rating      int    `json:"rating" bindig:"required" db:"rating"`
}

type ActorWithFilms struct {
	Actor
	Films []string `db:"films_list"`
}
