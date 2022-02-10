package db

import (
	"go-todo/users"
	"gorm.io/gorm"
)

func runMigration(db *gorm.DB) {
	db.AutoMigrate(&users.User{})

}
