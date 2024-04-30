package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
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

func NewUserModel(ID uuid.UUID, nome, email string, isActive bool, createdAt time.Time) *UserModel {
	return &UserModel{
		ID:        ID,
		Nome:      nome,
		Email:     email,
		IsActive:  isActive,
		CreatedAt: createdAt,
	}
}

func (this UserModel) Validate() bool {
	return this.isIDEmpty() || this.isNomeEmpty() || this.isEmailEmpty()
}

func (this UserModel) isIDEmpty() bool {
	return this.ID == uuid.Nil
}

func (this UserModel) isNomeEmpty() bool {
	return this.Nome == ""
}

func (this UserModel) isEmailEmpty() bool {
	return this.Email == ""
}

// func (this *User) IsUserTypeValid() bool {
// 	switch this.UserType {
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
