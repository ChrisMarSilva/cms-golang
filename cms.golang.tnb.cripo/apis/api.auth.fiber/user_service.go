package main

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
)

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (h UserService) Login(c *fiber.Ctx, payload UserRequest) (string, error) {
	user, err := h.userRepo.GetByEmail(c.Context(), payload.Email)
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

	sess, err := store.Get(c)
	if err != nil {
		log.Error("Erro no store: ", err.Error())
		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	sess.Set("jwt", tokenStr)
	if err := sess.Save(); err != nil {
		log.Error("Erro no token: ", err.Error())
		return "", err // c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	sess.SetExpiry(time.Hour * 24)

	return tokenStr, err
}

func (h UserService) Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	sess.Destroy()

	if err := sess.Save(); err != nil {
		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return nil // c.SendStatus(fiber.StatusOK)
}

func (h UserService) Refresh(c *fiber.Ctx) (jwt.MapClaims, error) {
	sess, err := store.Get(c)
	if err != nil {
		return nil, err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return nil, err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token.") // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	//sess.SetExpiry(time.Second * 2)
	// if err := sess.Save(); err != nil {
	//     return nil, err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	// }

	return claims, nil
}

func (h UserService) Verify(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err // c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return err // c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	return nil
}
