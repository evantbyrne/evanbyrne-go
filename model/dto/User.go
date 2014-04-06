package dto

type User struct {
	Email string `db:"email"`
	Password string `db:"password"`
	Salt string `db:"salt"`
}

type UserSession struct {
	Secret string `db:"secret"`
	UserEmail string `db:"user_email"`
}