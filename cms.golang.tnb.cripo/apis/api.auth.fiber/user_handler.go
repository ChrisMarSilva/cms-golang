package main

import (
	"database/sql"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h UserHandler) Home(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	return c.Status(fiber.StatusOK).SendString("I'm a GET / request!")
}

func (h UserHandler) Login(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	payload := new(UserRequest)
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		log.Error("Erro no payload:", err.Error())
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	// log.Info("Payload: ", payload)

	token, err := h.service.Login(c.Context(), *payload)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("Erro no StatusBadRequest:", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		log.Error("Erro no service: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	sess, err := store.Get(c)
	if err != nil {
		log.Error("Erro no store: ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	sess.Set("jwt", token)
	if err := sess.Save(); err != nil {
		log.Error("Erro no token: ", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	//log.Info("Token: ", token)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func (h UserHandler) Logout(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	// err = h.service.Logout(c.Context())
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	// }

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	sess.Destroy()
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h UserHandler) Refresh(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	// err = h.service.Refresh(c.Context())
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	// }

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return c.Status(fiber.StatusOK).JSON(claims)
	}

	//sess.SetExpiry(time.Second * 2)
	// if err := sess.Save(); err != nil {
	//     return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	// }

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
}

func (h UserHandler) Verify(c fiber.Ctx) error {
	log.Info(c.Method(), " ", c.Path())

	// err = h.service.Verify(c.Context())
	// if err != nil {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	// }

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	tokenStr := sess.Get("jwt")
	if tokenStr == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token found"})
	}

	token, err := jwt.Parse(tokenStr.(string), func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	return c.SendStatus(fiber.StatusOK)
}
