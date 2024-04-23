package main

import (
	"database/sql"

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
	return c.Status(fiber.StatusOK).SendString("I'm a GET / request!")
}

func (h UserHandler) Register(c *fiber.Ctx) error {
	payload := new(UserRegisterRequest)
	if err := json.Unmarshal(c.Body(), &payload); err != nil {
		log.Error("Erro no payload:", err.Error())
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}
	// log.Info("Payload: ", payload)

	user, err := h.service.Register(c, *payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	response := fiber.Map{"status": "success", "data": fiber.Map{"user": user}}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h UserHandler) Login(c *fiber.Ctx) error {

	// Validate user input (username/email, password)
	// Retrieve user data from the database based on input
	// Compare hashed password with input password
	// Generate a session or token for authentication
	// Return a success message or error response

	payload := new(UserLoginRequest)
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
	response := UserLoginResponse{Token: token}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h UserHandler) Logout(c *fiber.Ctx) error {
	err := h.service.Logout(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h UserHandler) Refresh(c *fiber.Ctx) error {
	response, err := h.service.Refresh(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h UserHandler) Verify(c *fiber.Ctx) error {
	err := h.service.Verify(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
