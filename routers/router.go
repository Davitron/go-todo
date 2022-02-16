package routers

import (
	"github.com/gin-gonic/gin"
	"go-todo/routers/api"
)

func InitRourter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/users", api.CreateUser)
	r.POST("/users/login", api.Auth)
	return r
}
