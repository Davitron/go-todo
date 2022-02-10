package api

import (
	"github.com/gin-gonic/gin"
	"go-todo/config"
	user "go-todo/users"
	"gorm.io/gorm"
)

var users user.UserService

func loadRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Configurations) {
	users.DB = db
	users.Config = cfg
	r.POST("/users", users.CreateUser)
	r.POST("/users/login", users.Authenticate)
	//User Routes

	//Task Routes
}
