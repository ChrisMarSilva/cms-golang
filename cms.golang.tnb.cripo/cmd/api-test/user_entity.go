package main

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID uuid.UUID `db:"id"`
	//Role     UsersRole`
	//Username   string    `db:"username"`
	Nome     string `db:"nome"`
	Email    string `db:"email"`
	Password string `db:"password"`
	//Avatar Photo     string    `db:"avatar"`
	//IsAdmin    bool      `db:"is_admin"`
	//IsBlocked  bool      `db:"is_blocked"`
	//IsVerified bool      `jdb:"is_verified"`
	IsActive  bool      `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	//UpdatedAt time.Time `db:"updated_at"`
	//DeletedAt time.Time `jsdb:"deleted_at"`
}

func NewUserEntity(ID uuid.UUID, nome, email string, isActive bool, createdAt time.Time) *UserEntity {
	return &UserEntity{
		ID:        ID,
		Nome:      nome,
		Email:     email,
		IsActive:  isActive,
		CreatedAt: createdAt,
	}
}

func (u UserEntity) Validate() bool {
	return u.isIDEmpty() || u.isNomeEmpty() || u.isEmailEmpty()
}

func (u UserEntity) isIDEmpty() bool {
	return u.ID == uuid.Nil
}

func (u UserEntity) isNomeEmpty() bool {
	return u.Nome == ""
}

func (u UserEntity) isEmailEmpty() bool {
	return u.Email == ""
}

// func (u *User) IsUserTypeValid() bool {
// 	switch u.UserType {
// 	case auth.CorporateUser:
// 		fallthrough
// 	case auth.Admin:
// 		fallthrough
// 	case auth.IndividualUser:
// 		return true
// 	default:
// 		return false
// 	}
// }
