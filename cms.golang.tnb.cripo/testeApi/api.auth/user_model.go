package main

import (
	"time"

	"github.com/google/uuid"
)

// User is the model for a user.

type UserModel struct {
	ID uuid.UUID `json:"id" db:"id"`
	//Username   string    `json:"username" db:"username"`
	Nome     string `json:"nome" db:"nome"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	//Avatar     string    `json:"avatar" db:"avatar"`
	//IsAdmin    bool      `json:"is_admin" db:"is_admin"`
	//IsBlocked  bool      `json:"is_blocked" db:"is_blocked"`
	//IsVerified bool      `json:"is_verified" db:"is_verified"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	//Updated_at time.Time `json:"updated_at" db:"updated_at"`
	//Deleted_at time.Time `json:"deleted_at" db:"deleted_at"`
	//LastLogin  time.Time `json:"last_login" db:"last_login"`
}
