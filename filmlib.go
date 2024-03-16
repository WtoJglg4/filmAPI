package filmapi

type Actor struct {
	Id        int    `json:"-"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
}

type Film struct {
	Id          int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      int    `json:"rating"`
}
