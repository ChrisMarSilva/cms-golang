package main

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (h UserService) Login(ctx context.Context, payload UserRequest) (string, error) {
	user, err := h.userRepo.GetByEmail(ctx, payload.Email)
	if err != nil {
		log.Error("Erro no repository:", err.Error())
		return "", err
	}

	if user.Password != payload.Password {
		log.Error("Erro no Password: Senha inválida.")
		return "", errors.New("Senha inválida.")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":        user.ID,
		"nome":      user.Nome,
		"email":     user.Email,
		"is_active": user.IsActive,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secretKey))
	return tokenStr, err
}

func (h UserService) Logout(ctx context.Context) error {
	return nil
}

func (h UserService) Refresh(ctx context.Context) error {
	return nil
}

func (h UserService) Verify(ctx context.Context) error {
	return nil
}
