package main

import (
	"time"

	"github.com/google/uuid"
)

type UsersRole string

const (
	AdminRole UsersRole = "admin"
	UserRole  UsersRole = "user"
	GuestRole UsersRole = "guest"
)

type UserModel struct {
	ID uuid.UUID `json:"id" db:"id"`
	//Username   string    `json:"username" db:"username"`
	Nome     string `json:"nome" db:"nome"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	//Avatar Photo     string    `json:"avatar" db:"avatar"`
	//IsAdmin    bool      `json:"is_admin" db:"is_admin"`
	//IsBlocked  bool      `json:"is_blocked" db:"is_blocked"`
	//IsVerified bool      `json:"is_verified" db:"is_verified"`
	//UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	//DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
	//LastLogin  time.Time `json:"last_login" db:"last_login"
	//Role     UsersRole`
	IsActive  bool      `json:"-" db:"is_active"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

func NewUser(ID uuid.UUID, nome string, email string) *UserModel {
	return &UserModel{
		ID:    ID,
		Nome:  nome,
		Email: email,
	}
}

type UserRegisterRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	//PasswordConfirm string `json:"passwordConfirm""`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type UserLoginRequest struct {
	// LoginRequest
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	// LoginResponse
	Token string `json:"token"`
}