package handlers

import (
	"encoding/json"
	"fmt"
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

	// idStr := strings.Trim(ctx.Params("id"), " ")
	// if idStr == "" {
	// 	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Id está vazio."})
	// }

	// clienteId, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	// }

	clienteId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		//return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"message": "O Cliente ID deve ser um número inteiro.",
	}

	payload := new(dtos.TransacaoRequestDto)

	if err := json.Unmarshal(ctx.Body(), &payload); err != nil { // if err := ctx.BodyParser(payload); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Payload inválido: " + err.Error()})
	}

	payload.Tipo = strings.ToLower(strings.Trim(payload.Tipo, " "))
	payload.Descricao = strings.Trim(payload.Descricao, " ")

	if err := payload.Valido(); err != nil {
		//ErroValorDaTransacao
		//  ErroTipoDaTransacao
		//ErroDescricao
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Payload inválido: " + err.Error()})
	}

	response, err := h.service.CreateTransaction(clienteId, *payload)
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

	// strId := strings.Trim(ctx.Params("id"), " ")
	// if strId == "" {
	// 	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Id está vazio."})
	// }

	// clienteId, err := strconv.Atoi(strId)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
	// }

	clienteId, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": err.Error()})
		//return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"message": "O Cliente ID deve ser um número inteiro.",
	}

	response, err := h.service.GetExtract(clienteId)
	if err != nil {
		status := fiber.StatusUnprocessableEntity
		if strings.Contains(fmt.Sprint(err), "Cliente não localizado.") {
			status = fiber.StatusNotFound
			//return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"message": "Cliente não existe."})
		}
		return ctx.Status(status).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
