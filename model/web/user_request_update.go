package web

type UserRequestUpdate struct {
	Id       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone"`
}
