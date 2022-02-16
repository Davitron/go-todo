package models

import (
	"log"
)

type Users struct {
	Model
	Username string `json:"username"`
	Password string `json:"password" `
	Email    string `json:"email"`
}

func UserExistsByEmail(email string) (*Users, bool) {
	var user *Users
	query := db.Select("users.*")
	query = query.Group("users.id")
	if rows := query.Where("users.email = ?", email).First(&user).RowsAffected; rows > 0 {
		return user, true
	}
	return user, false
}

func CreateUser(data map[string]interface{}) error {
	user := &Users{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Failed to create user ", err)
		return err
	}
	return nil
}
