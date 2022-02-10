package users

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password" `
	Email    string `json:"email"`
}
