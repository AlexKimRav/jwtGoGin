package controller

import (
	"jwtgogin/data/request"
	"jwtgogin/data/response"
	"jwtgogin/helper"
	"jwtgogin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	AuthenticationService service.AuthenticationService
}

func NewAutheticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{AuthenticationService: service}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	LoginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&LoginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.AuthenticationService.Login(LoginRequest)

	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invlid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successefully log in!",
		Data:    resp,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	helper.ErrorPanic(err)

	controller.AuthenticationService.Register(createUsersRequest)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successefully log in!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
