package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func GenerateJWT(user UserModel) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"nome":      user.Nome,
		"email":     user.Email,
		"is_active": user.IsActive,
		//"role":      user.Role,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	//tokenStr, err := token.SignedString([]byte(SecretKey))
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
