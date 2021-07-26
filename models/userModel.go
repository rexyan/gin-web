package models

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Gender   uint8  `db:"gender"`
}
