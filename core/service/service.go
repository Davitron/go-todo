package service

import (
	"go-todo/config"
	"gorm.io/gorm"
)

type Service struct {
	DB     *gorm.DB
	Config *config.Configurations
}
