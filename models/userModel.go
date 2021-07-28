package models

type User struct {
	UserID   int64  `db:"user_id" json:"user_id"`
	UserName string `db:"username" json:"user_name"`
	Password string `db:"password" json:"-"`
	Email    string `db:"email" json:"email"`
	Gender   uint8  `db:"gender" json:"gender"`
}
