package controller

import (
	"crud-mysql/helper"
	"crud-mysql/model/web"
	"crud-mysql/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImplementation{
		UserService: userService,
	}
}

func (controller *UserControllerImplementation) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	userRequest := web.UserRequest{}
	helper.ReadFromRequestBody(request, &userRequest)

	userResponse := controller.UserService.Create(request.Context(), userRequest)
	webResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *UserControllerImplementation) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserRequestUpdate{}
	helper.ReadFromRequestBody(request, &userRequest)

	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userRequest)
	webResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImplementation) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *UserControllerImplementation) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImplementation) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := controller.UserService.FindAll(request.Context())
	fmt.Println(userResponse)
	webResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImplementation) ChangePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRequest := web.UserRequestUpdatePassword{}
	helper.ReadFromRequestBody(request, &userRequest)

	userId := params.ByName("id")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userRequest.Id = id
	userResponse := controller.UserService.ChangePassword(request.Context(), userRequest)
	webResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
