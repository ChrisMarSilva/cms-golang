package main

type UserRegisterRequest struct {
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//Username string `json:"username" binding:"required"`
	//PasswordConfirm string `json:"passwordConfirm""`
}
