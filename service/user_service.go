package service

import (
	"context"
	"crud-mysql/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserRequest) web.UserResponse
	Update(ctx context.Context, request web.UserRequestUpdate) web.UserResponse
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []web.UserResponse
	FindById(ctx context.Context, id int) web.UserResponse
	ChangePassword(ctx context.Context, request web.UserRequestUpdatePassword) web.UserResponse
}
