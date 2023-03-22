package web

type UserRequestUpdatePassword struct {
	Id       int    `json:"id" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
