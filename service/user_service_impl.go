package service

import (
	"context"
	"crud-mysql/helper"
	"crud-mysql/model/domain"
	"crud-mysql/model/web"
	"crud-mysql/repository"
	"database/sql"

	"github.com/go-playground/validator"
)

type UserServiceImplementation struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImplementation{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImplementation) Create(ctx context.Context, request web.UserRequest) web.UserResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
	}

	service.UserRepository.Create(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImplementation) Update(ctx context.Context, request web.UserRequestUpdate) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Name = request.Name
	user.Username = request.Username
	user.Email = request.Email
	user.Phone = request.Phone

	service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImplementation) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImplementation) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)
}

func (service *UserServiceImplementation) FindById(ctx context.Context, id int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImplementation) ChangePassword(ctx context.Context, request web.UserRequestUpdatePassword) web.UserResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	user.Password = request.Password

	service.UserRepository.ChangePassword(ctx, tx, user)

	return helper.ToUserResponse(user)

}
