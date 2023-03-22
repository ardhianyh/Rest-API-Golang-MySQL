package app

import (
	"crud-mysql/controller"
	"crud-mysql/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/v1/users", userController.FindAll)
	router.GET("/v1/users/:id", userController.FindById)
	router.POST("/v1/users", userController.Create)
	router.PUT("/v1/users/:id", userController.Update)
	router.PUT("/v1/users/:id/change-password", userController.ChangePassword)
	router.DELETE("/v1/users/:id", userController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
