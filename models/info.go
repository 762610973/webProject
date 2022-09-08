package models

type Info struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Gender   int8   `json:"gender" db:"gender"`
	UserID   int64  `db:"author_id"`
}
