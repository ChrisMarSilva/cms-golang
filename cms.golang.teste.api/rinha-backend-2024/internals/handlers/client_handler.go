package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/rinha-backend-2024/internals/dtos"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/services"
	"github.com/gofiber/fiber/v3"
)

type IClientHandler interface {
	CreateTransaction(ctx *fiber.Ctx) error
	CreateTransactionBatch(ctx *fiber.Ctx, batchCreateTransaction chan dtos.Job) error
	GetExtract(ctx *fiber.Ctx) error
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

	//var payload dtos.TransacaoRequestDto
	payload := new(dtos.TransacaoRequestDto)

	if err := sonic.Unmarshal(ctx.Body(), &payload); err != nil { // if err := ctx.BodyParser(payload); err != nil {
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

	response, err := h.service.CreateTransaction(ctx.Context(), clienteId, *payload)
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

	response, err := h.service.GetExtract(ctx.Context(), clienteId)
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

func (h *ClientHandler) CreateTransactionBatch(ctx fiber.Ctx, batchCreateTransaction chan models.Job) error {
	clienteId, _ := ctx.ParamsInt("id")
	parseTransactionChannel := make(chan models.ParseTransactionResult)
	go parseTransaction(clienteId, ctx.Body(), parseTransactionChannel)

	parseResult := <-parseTransactionChannel
	if parseResult.Error != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Payload inválido: " + parseResult.Error.Error()})
	}

	transacao := parseResult.Transacao
	// transacao = dtos.TransacaoResponseDto{
	// 	Limite: cliente.Limite,
	// 	Saldo:  cliente.Saldo,
	// }

	//go persistTransaction(batchCreateTransaction, transacao)
	go func() {
		batchCreateTransaction <- models.Job{Name: "", Transacao: transacao}
	}()

	return ctx.Status(fiber.StatusOK).JSON("response")
}

func parseTransaction(clienteId int, bytes []byte, result chan models.ParseTransactionResult) {
	payload := new(dtos.TransacaoRequestDto)
	if err := sonic.Unmarshal(bytes, &payload); err != nil {
		result <- models.ParseTransactionResult{
			Transacao: nil,
			Error:     err,
		}
	}

	payload.Tipo = strings.ToLower(strings.Trim(payload.Tipo, " "))
	payload.Descricao = strings.Trim(payload.Descricao, " ")
	if err := payload.Valido(); err != nil {
		result <- models.ParseTransactionResult{Transacao: nil, Error: fmt.Errorf("field cannot be null")}
	}

	transacao := models.ClienteTransacao{
		IdCliente: clienteId,
		Valor:     payload.Valor,
		Tipo:      payload.Tipo,
		Descricao: payload.Descricao,
	}

	result <- models.ParseTransactionResult{Transacao: &transacao, Error: nil}
}

func (h *ClientHandler) ProcessBatch(batchChannel chan models.Job, batchHandler BatchHandler) {
	batchSleep := 1000
	batchSize := 2000

	for {
		var batchJobs []models.Job
		batchJobs = append(batchJobs, <-batchChannel)
		time.Sleep(time.Duration(batchSleep) * time.Millisecond)

		for i := 0; i < batchSize; i++ {
			select {
			case job := <-batchChannel:
				batchJobs = append(batchJobs, job)
			default:
				break
			}
		}

		batchHandler(batchJobs)
	}
}
