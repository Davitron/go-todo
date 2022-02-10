package api

import (
	"github.com/gin-gonic/gin"
	"go-todo/config"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(db *gorm.DB, cfg *config.Configurations) *Server {
	server := &Server{}
	server.setupRouter(db, cfg)
	return server
}

func (server *Server) setupRouter(db *gorm.DB, cfg *config.Configurations) {
	router := gin.Default()
	server.db = db
	loadRoutes(router, db, cfg)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
