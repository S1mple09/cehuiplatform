package models

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
	Avatar   string `db:"avatar"`
}
