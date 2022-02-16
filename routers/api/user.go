package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-todo/dto"
	"go-todo/pkg/settings"
	"go-todo/services"
	"go-todo/utils"
	"net/http"
)

var userReq *dto.NewUserRequestDto
var userRes *dto.NewUserResponseDto

var authReq *dto.UserAuthRequestDto
var authRes *dto.UserAuthResponseDto

func CreateUser(ctx *gin.Context) {
	var err error
	if err = ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	userService := &services.UserService{
		Email:    userReq.Email,
		Password: userReq.Password,
		Username: userReq.Username,
	}

	if _, exists := userService.Exists(userReq.Email); exists {
		ctx.JSON(
			http.StatusConflict,
			utils.ErrorResponse(errors.New("user already exists")))
		return
	}

	if err = userService.Create(); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			utils.ErrorResponse(errors.New("failed to create new user")))
		return
	}

	userRes = &dto.NewUserResponseDto{
		Email:    userService.Email,
		Username: userService.Username,
	}

	ctx.JSON(http.StatusOK, userRes)
	return
}

func Auth(ctx *gin.Context) {
	var err error
	if err = ctx.ShouldBindJSON(&authReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	userService := &services.UserService{}
	user, exists := userService.Exists(authReq.Email)
	err = utils.CheckPassword(authReq.Password, user.Password)
	if !exists || err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			utils.ErrorResponse(errors.New("invalid login details")))
		return
	}

	token, err := utils.GenerateToken(user.Username, user.Email, settings.AppSettings.JWT.SecretKey)
	authRes = &dto.UserAuthResponseDto{
		Message: "Authentication successful",
		Token:   token,
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(authRes))
	return
}
