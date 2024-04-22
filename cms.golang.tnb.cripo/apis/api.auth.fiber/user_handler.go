package main

import (
	"database/sql"
	log2 "log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h UserHandler) Home(c *fiber.Ctx) error {
	log.Info("request /")
	log.Debug("Are you OK?")
	log2.Println("request /")

	return c.Status(fiber.StatusOK).SendString("I'm a GET / request!")
}

func (h UserHandler) Login(c *fiber.Ctx) error {
	payload := new(UserRequest)
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		log.Error("Erro no payload:", err.Error())
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}
	// log.Info("Payload: ", payload)

	token, err := h.service.Login(c, *payload)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("Erro no StatusBadRequest:", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		log.Error("Erro no service: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	//log.Info("Token: ", token)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func (h UserHandler) Logout(c *fiber.Ctx) error {
	err := h.service.Logout(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h UserHandler) Refresh(c *fiber.Ctx) error {
	claims, err := h.service.Refresh(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(claims)
}

func (h UserHandler) Verify(c *fiber.Ctx) error {
	err := h.service.Verify(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
