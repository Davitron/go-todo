package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-todo/core/service"
	"go-todo/utils"
	"net/http"
)

type UserService struct {
	service.Service
}

var userRepository UserRepository

func (userService *UserService) CreateUser(ctx *gin.Context) {
	var req NewUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if isConflict := userRepository.IsExisting(req.Email, userService.DB); isConflict {
		ctx.JSON(http.StatusConflict, utils.SuccessResponse("User already exists"))
		return
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	err = userRepository.Create(&req, hashedPassword, userService.DB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.SuccessResponse("User Created"))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse("User Created"))
	return
}

func (userService *UserService) Authenticate(ctx *gin.Context) {
	var req UserAuthRequest
	var res *UserAuthResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if exists := userRepository.IsExisting(req.Email, userService.DB); exists == false {
		ctx.JSON(http.StatusNotFound, utils.SuccessResponse("User does not exist"))
		return
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(errors.New("invalid authentication credentials")))
		return
	}
	token, err := utils.GenerateToken(user.Username, user.Email, userService.Config.JWT.SecretKey)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			utils.ErrorResponse(errors.New("sorry! unable to authenticate at this point")))
		return
	}
	res = &UserAuthResponse{
		Message: "authentication successful",
		Token:   token,
	}

	ctx.JSON(http.StatusUnauthorized, utils.SuccessResponse(res))
	return
}
