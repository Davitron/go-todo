package users

type NewUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type UserAuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserAuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type NewUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
}
