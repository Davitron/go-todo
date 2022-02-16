package dto

type NewUserRequestDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type NewUserResponseDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserAuthRequestDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserAuthResponseDto struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
