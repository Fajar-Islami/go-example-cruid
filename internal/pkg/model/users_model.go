package model

type CreateUser struct {
	Email    string `json:"title" validate:"required,email"`
	Password string `json:"description" validate:"required"`
	Name     string `json:"author" validate:"required"`
}

type Login struct {
	Email    string `json:"title" validate:"required,email"`
	Password string `json:"description" validate:"required"`
}

type LoginRes struct {
	Email string `json:"title"`
	Name  string `json:"author"`
	Token string `json:"token"`
}
