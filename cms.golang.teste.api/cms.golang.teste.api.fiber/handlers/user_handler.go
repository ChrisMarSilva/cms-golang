package handlers

import (
	"errors"
	"fmt"
	"strings"

	entity "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/entities"
	repository "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories"
	service "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	opentracing "github.com/opentracing/opentracing-go"
	spanlog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

type UserHandler struct {
	service service.UserService
	logRepo repository.LogMonitorRepositoryMSSQL
}

func NewUserHandler(service service.UserService, logRepo repository.LogMonitorRepositoryMSSQL) *UserHandler {
	return &UserHandler{
		service: service,
		logRepo: logRepo,
	}
}

func (handler *UserHandler) GetAll(c *fiber.Ctx) error {

	metodo := "UserHandler.GetAll"

	// ctx := context.Background() // ctx := c.Context()

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), metodo)
	sp.SetTag("Method", c.Method())
	sp.SetTag("URL", c.Path())
	defer sp.Finish()

	users, err := handler.service.GetAll(ctx)
	if err != nil {
		handler.logRepo.Inserir(metodo, "1", "URL: "+c.Path()+"; Erro: "+err.Error())
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	sp.SetTag("LEN", len(users))

	return c.Status(fiber.StatusOK).JSON(entity.ResponseHTTP{Success: true, Message: "Success get all users.", Data: users, Length: len(users)})
}

func (handler *UserHandler) Get(c *fiber.Ctx) error {

	metodo := "UserHandler.Get"

	// ctx := context.Background() // ctx := c.Context()

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), metodo)
	sp.SetTag("Method", c.Method())
	sp.SetTag("URL", c.Path())
	defer sp.Finish()

	tracer := opentracing.GlobalTracer()

	// id, err := c.ParamsInt("id")
	// id, _ := strconv.Atoi(c.Params("id"))
	sp2 := tracer.StartSpan("UserHandler.Params.id", opentracing.ChildOf(sp.Context()))
	id := c.Params("id")
	if id == "" {
		sp2.SetTag("error", true)
		sp2.LogFields(spanlog.Error(errors.New("id empty")))
		sp2.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: "id empty", Data: nil})
	}
	sp2.Finish()

	sp.SetTag("id", id)

	sp3 := tracer.StartSpan("UserHandler.uuid.Parse", opentracing.ChildOf(sp.Context()))
	idParse, err := uuid.Parse(id)
	if err != nil {
		sp3.SetTag("error", true)
		sp3.LogFields(spanlog.Error(err))
		sp3.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}
	sp3.Finish()

	user, err := handler.service.Get(ctx, idParse)
	if err != nil {
		handler.logRepo.Inserir(metodo, "1", "URL: "+c.Path()+"; Erro: "+err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: fmt.Sprintf("User with ID %v not found.", id), Data: nil})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	return c.Status(fiber.StatusOK).JSON(entity.ResponseHTTP{Success: true, Message: "Success get a user.", Data: user})
}

func (handler *UserHandler) Create(c *fiber.Ctx) error {

	metodo := "UserHandler.Create"

	// ctx := context.Background() // ctx := c.Context()

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), metodo)
	sp.SetTag("Method", c.Method())
	sp.SetTag("URL", c.Path())
	defer sp.Finish()

	tracer := opentracing.GlobalTracer()

	sp2 := tracer.StartSpan("UserHandler.BodyParser", opentracing.ChildOf(sp.Context()))
	var body entity.User
	err := c.BodyParser(&body)
	if err != nil {
		sp2.SetTag("error", true)
		sp2.LogFields(spanlog.Error(err))
		sp2.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil}) // "cannot parse json"
	}
	sp2.Finish()

	sp.SetTag("body", body)

	user := entity.User{ID: uuid.New(), Nome: strings.Trim(body.Nome, " "), Status: entity.UserActive}

	err = handler.service.Create(ctx, user)
	if err != nil {
		handler.logRepo.Inserir(metodo, "1", "URL: "+c.Path()+"; Erro: "+err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	return c.Status(fiber.StatusOK).JSON(entity.ResponseHTTP{Success: true, Message: "Success create a user.", Data: user})
}

func (handler *UserHandler) Update(c *fiber.Ctx) error {

	metodo := "UserHandler.Update"

	// ctx := context.Background() // ctx := c.Context()

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), metodo)
	sp.SetTag("Method", c.Method())
	sp.SetTag("URL", c.Path())
	defer sp.Finish()

	tracer := opentracing.GlobalTracer()

	sp2 := tracer.StartSpan("UserHandler.Params.id", opentracing.ChildOf(sp.Context()))
	id := c.Params("id")
	if id == "" {
		sp2.SetTag("error", true)
		sp2.LogFields(spanlog.Error(errors.New("id empty")))
		sp2.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: "id empty", Data: nil})
	}
	sp2.Finish()

	sp.SetTag("id", id)

	sp3 := tracer.StartSpan("UserHandler.uuid.Parse", opentracing.ChildOf(sp.Context()))
	idParse, err := uuid.Parse(id)
	if err != nil {
		sp3.SetTag("error", true)
		sp3.LogFields(spanlog.Error(err))
		sp3.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}
	sp3.Finish()

	sp4 := tracer.StartSpan("UserHandler.BodyParser", opentracing.ChildOf(sp.Context()))
	var body entity.User
	err = c.BodyParser(&body)
	if err != nil {
		sp4.SetTag("error", true)
		sp4.LogFields(spanlog.Error(err))
		sp4.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}
	sp4.Finish()

	err = handler.service.Update(ctx, idParse, body)
	if err != nil {
		handler.logRepo.Inserir(metodo, "1", "URL: "+c.Path()+"; Erro: "+err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	return c.Status(fiber.StatusOK).JSON(entity.ResponseHTTP{Success: true, Message: "Success update user.", Data: body})
}

func (handler *UserHandler) Delete(c *fiber.Ctx) error {

	metodo := "UserHandler.Delete"

	// ctx := context.Background() // ctx := c.Context()

	sp, ctx := opentracing.StartSpanFromContext(c.Context(), metodo)
	sp.SetTag("Method", c.Method())
	sp.SetTag("URL", c.Path())
	defer sp.Finish()

	tracer := opentracing.GlobalTracer()

	sp2 := tracer.StartSpan("UserHandler.Params.id", opentracing.ChildOf(sp.Context()))
	id := c.Params("id")
	if id == "" {
		sp2.SetTag("error", true)
		sp2.LogFields(spanlog.Error(errors.New("id empty")))
		sp2.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: "id empty", Data: nil})
	}
	sp2.Finish()

	sp.SetTag("id", id)

	sp3 := tracer.StartSpan("UserHandler.uuid.Parse", opentracing.ChildOf(sp.Context()))
	idParse, err := uuid.Parse(id)
	if err != nil {
		sp3.SetTag("error", true)
		sp3.LogFields(spanlog.Error(err))
		sp3.Finish()
		return c.Status(fiber.StatusBadRequest).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}
	sp3.Finish()

	err = handler.service.Delete(ctx, idParse)
	if err != nil {
		handler.logRepo.Inserir(metodo, "1", "URL: "+c.Path()+"; Erro: "+err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(entity.ResponseHTTP{Success: false, Message: err.Error(), Data: nil})
	}

	return c.Status(fiber.StatusOK).JSON(entity.ResponseHTTP{Success: true, Message: "Success delete user.", Data: nil})
}
