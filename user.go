package filmapi

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" bindig:"required" db:"username"`
	Password string `json:"password" bindig:"required" db:"password_hash"`
	Role     string `json:"-" bindig:"required" db:"role"`
}
