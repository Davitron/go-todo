package services

import (
	"go-todo/models"
	"go-todo/utils"
)

type UserService struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *UserService) Create() error {
	var err error
	hashedpassword, err := utils.HashPassword(u.Password)
	user := map[string]interface{}{
		"username": u.Username,
		"password": hashedpassword,
		"email":    u.Email,
	}
	if err = models.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (u *UserService) Exists(email string) (*models.Users, bool) {
	user, exists := models.UserExistsByEmail(email)
	return user, exists
}
