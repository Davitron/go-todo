package users

import (
	"gorm.io/gorm"
	"log"
)

type UserRepository struct{}

var user *User

func (u *UserRepository) IsExisting(email string, db *gorm.DB) bool {
	query := db.Select("users.*")
	query = query.Group("users.id")
	if rows := query.Where("users.email = ?", email).First(&user).RowsAffected; rows > 0 {
		return true
	}
	return false
}

func (u *UserRepository) Create(req *NewUserRequest, hashedPassword string, db *gorm.DB) error {
	newuser := &User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	if err := db.Create(&newuser).Error; err != nil {
		log.Fatal("Failed to create user ", err)
		return err
	}
	return nil
}
