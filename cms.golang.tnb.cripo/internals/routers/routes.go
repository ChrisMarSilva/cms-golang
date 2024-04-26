package routers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/golang-jwt/jwt/v5"
)

func ConfigRoutes(app *fiber.App) *fiber.App {
	//Database
	db, err := GetDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %s", err.Error())
	}
	// defer db.Close()

	//Repository
	userRepo := NewUserRepository(db)

	//Service
	userServ := NewUserService(*userRepo)

	//Handle
	userHandler := NewUserHandler(*userServ)



	
	dbRepo := repository.NewDBRepo(db.SQL)
	userRepo := repository.NewUserRepo(db.SQL)
	router := chi.NewRouter()

	
	//routes

	app.Use(AuthMiddleware)

	jwt := NewAuthMiddleware(SecretKey)

	app.Get("/", timeout.New(userHandler.Home, 5*time.Second))
	app.Get("/accessible", accessible)
	app.Get("/restricted", jwt, restricted)
	// app.Get("/users/me", middleware.DeserializeUser, controllers.GetMe)
	// app.Get("/metrics", monitor.New())

	routes := app.Group("/api/v1/auth")
	routes.Post("/login", userHandler.Login)
	routes.Post("/logout", userHandler.Logout)
	routes.Get("/refresh", userHandler.Refresh)
	routes.Get("/verify", userHandler.Verify)

	app.Use(func(c *fiber.Ctx) error {
		if err := recover(); err != nil {
			// Handle the error and respond with an error message
			return c.Status(500).SendString("Internal Server Error")
		}
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	return app
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	name := claims["nome"].(string)
	email := claims["email"].(string)
	return c.SendString("Welcome " + name + " " + email + " " + id)
}



package routes

import (
    "github.com/Siddheshk02/jwt-auth-api/controllers" // importing the routes package 
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    api := app.Group("/user")
    api.Get("/get-user", controllers.User)

    api.Post("/register", controllers.Register)

    api.Post("/login", controllers.Login)

    api.Post("/logout", controllers.Logout)
}
