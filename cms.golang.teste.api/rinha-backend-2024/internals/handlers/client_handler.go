package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/chrismarsilva/rinha-backend-2024/internals/dtos"
	"github.com/chrismarsilva/rinha-backend-2024/internals/services"
	"github.com/gofiber/fiber/v3"
)

type IClientHandler interface {
	CreateTransaction(c *fiber.Ctx) error
	GetExtract(c *fiber.Ctx) error
}

type ClientHandler struct {
	service services.ClientService
}

func NewClientHandler(service services.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

func (h *ClientHandler) CreateTransaction(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	idStr := strings.Trim(ctx.Params("id"), " ")
	if idStr == "" {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Id está vazio."})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	payload := new(dtos.TransacaoRequestDto)

	if err := json.Unmarshal(ctx.Body(), &payload); err != nil { // if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Payload inválido: " + err.Error()})
	}

	payload.Tipo = strings.ToLower(strings.Trim(payload.Tipo, " "))
	payload.Descricao = strings.Trim(payload.Descricao, " ")

	if err := payload.Valido(); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Payload inválido: " + err.Error()})
	}

	response, err := h.service.CreateTransaction(id, *payload)
	if err != nil {
		status := fiber.StatusUnprocessableEntity
		if strings.Contains(fmt.Sprint(err), "Cliente não localizado.") {
			status = fiber.StatusNotFound
		}
		return ctx.Status(status).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *ClientHandler) GetExtract(ctx fiber.Ctx) error {
	ctx.Accepts("application/json")

	strId := strings.Trim(ctx.Params("id"), " ")
	if strId == "" {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Id está vazio."})
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	}

	response, err := h.service.GetExtract(id)
	if err != nil {
		status := fiber.StatusUnprocessableEntity
		if strings.Contains(fmt.Sprint(err), "Cliente não localizado.") {
			status = fiber.StatusNotFound
		}
		return ctx.Status(status).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
