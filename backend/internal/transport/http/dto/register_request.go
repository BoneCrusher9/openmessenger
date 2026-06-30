package dto

type RegisterRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=32"`
	DisplayName string `json:"display_name" binding:"required,min=1,max=64"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8,max=72"`
}
