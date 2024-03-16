package filmapi

type Actor struct {
	Id        int    `json:"-"`
	Name      string `json:"name" bindig:"required"`
	Gender    string `json:"gender" bindig:"required"`
	BirthDate string `json:"birth_date" bindig:"required"`
}

type Film struct {
	Id          int    `json:"-"`
	Name        string `json:"name" bindig:"required"`
	Description string `json:"description" bindig:"required"`
	ReleaseDate string `json:"release_date" bindig:"required"`
	Rating      int    `json:"rating" bindig:"required"`
}
