package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Handler struct {
	routes map[string]*config.Route
	logger *zap.Logger
	db     *database.Database
}

func NewHandler(db *database.Database, logger *zap.Logger) *Handler {
	routes, err := db.GetRoutes()
	if err != nil {
		logger.Error("Failed to load routes", zap.Error(err))
	}

	routeMap := make(map[string]*config.Route)
	for _, route := range routes {
		routeMap[route.Path] = route
	}

	return &Handler{routes: routeMap, logger: logger, db: db}
}

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

type TaskHandler struct {
}

func (th TaskHandler) Routes(router *http.ServeMux) {
	router.HandleFunc("GET /tasks", th.getAllTasks)
	router.HandleFunc("GET /tasks/{id}", th.getTask)
	router.HandleFunc("POST /tasks", th.createTask)
	router.HandleFunc("PUT /tasks/{id}", th.updateTask)
	router.HandleFunc("DELETE /tasks/{id}", th.deleteTask)
	router.HandleFunc("POST /tasks/create-default", th.createDefaultTasks)

	router.HandleFunc("GET /tasks-xml", th.getAllTasksXml)
	router.HandleFunc("GET /tasks-xml/{id}", th.getTaskXml)

	router.Handle("GET /protected-apikey", ApiKeyAuthMiddleware(http.HandlerFunc(th.protectedApiKey)))
}

type UsersHandler interface {
	// ...
	Register(c *gin.Context)
	Login(c *gin.Context)
}

// Register register
// @Summary register
// @Description register
// @Tags auth
// @accept json
// @Produce json
// @Param data body types.RegisterRequest true "login information"
// @Success 200 {object} types.RegisterRespond{}
// @Router /api/v1/auth/register [post]
func (h *usersHandler) Register(c *gin.Context) {

}

// Login login
// @Summary login
// @Description login
// @Tags auth
// @accept json
// @Produce json
// @Param data body types.LoginRequest true "login information"
// @Success 200 {object} types.LoginRespond{}
// @Router /api/v1/teacher/login [post]
func (h *usersHandler) Login(c *gin.Context) {

}
