package main

import (
	"crud-mysql/app"
	"crud-mysql/controller"
	"crud-mysql/helper"
	"crud-mysql/repository"
	"crud-mysql/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	fmt.Println("Listen to port 8080")

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
