package routers

func AuthRoutes() {
	// r.POST("/login", controllers.Login)
	// r.POST("/signup", controllers.Signup)
	// r.GET("/home", controllers.Home)
	// r.GET("/premium", controllers.Premium)
	// r.GET("/logout", controllers.Logout)
}

// package routes

// import (
//     "github.com/gofiber/fiber/v2"

//     swagger "github.com/arsmn/fiber-swagger/v2"
// )

// // SwaggerRoute func for describe group of API Docs routes.
// func SwaggerRoute(a *fiber.App) {
//     // Create routes group.
//     route := a.Group("/swagger")

//     // Routes for GET method:
//     route.Get("*", swagger.Handler) // get one user by ID
// }

// // ./pkg/routes/not_found_route.go

// package routes

// import "github.com/gofiber/fiber/v2"

// // NotFoundRoute func for describe 404 Error route.
// func NotFoundRoute(a *fiber.App) {
//     // Register new special route.
//     a.Use(
//         // Anonimus function.
//         func(c *fiber.Ctx) error {
//             // Return HTTP 404 status and JSON response.
//             return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//                 "error": true,
//                 "msg":   "sorry, endpoint is not found",
//             })
//         },
//     )
// }

// // ./pkg/routes/private_routes.go

// package routes

// import (
//     "github.com/gofiber/fiber/v2"
//     "github.com/koddr/tutorial-go-fiber-rest-api/app/controllers"
//     "github.com/koddr/tutorial-go-fiber-rest-api/pkg/middleware"
// )

// // PrivateRoutes func for describe group of private routes.
// func PrivateRoutes(a *fiber.App) {
//     // Create routes group.
//     route := a.Group("/api/v1")

//     // Routes for POST method:
//     route.Post("/book", middleware.JWTProtected(), controllers.CreateBook) // create a new book

//     // Routes for PUT method:
//     route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

//     // Routes for DELETE method:
//     route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
// }

// // ./pkg/routes/private_routes.go

// package routes

// import (
//     "github.com/gofiber/fiber/v2"
//     "github.com/koddr/tutorial-go-fiber-rest-api/app/controllers"
// )

// // PublicRoutes func for describe group of public routes.
// func PublicRoutes(a *fiber.App) {
//     // Create routes group.
//     route := a.Group("/api/v1")

//     // Routes for GET method:
//     route.Get("/books", controllers.GetBooks)              // get list of all books
//     route.Get("/book/:id", controllers.GetBook)            // get one book by ID
//     route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
// }

// // ./main.go

// package main

// import (
//     "github.com/gofiber/fiber/v2"
//     "github.com/koddr/tutorial-go-fiber-rest-api/pkg/configs"
//     "github.com/koddr/tutorial-go-fiber-rest-api/pkg/middleware"
//     "github.com/koddr/tutorial-go-fiber-rest-api/pkg/routes"
//     "github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"

//     _ "github.com/joho/godotenv/autoload"                // load .env file automatically
//     _ "github.com/koddr/tutorial-go-fiber-rest-api/docs" // load API Docs files (Swagger)
// )

// // @title API
// // @version 1.0
// // @description This is an auto-generated API Docs.
// // @termsOfService http://swagger.io/terms/
// // @contact.name API Support
// // @contact.email your@mail.com
// // @license.name Apache 2.0
// // @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// // @securityDefinitions.apikey ApiKeyAuth
// // @in header
// // @name Authorization
// // @BasePath /api
// func main() {
//     // Define Fiber config.
//     config := configs.FiberConfig()

//     // Define a new Fiber app with config.
//     app := fiber.New(config)

//     // Middlewares.
//     middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

//     // Routes.
//     routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
//     routes.PublicRoutes(app)  // Register a public routes for app.
//     routes.PrivateRoutes(app) // Register a private routes for app.
//     routes.NotFoundRoute(app) // Register route for 404 Error.

//     // Start server (with graceful shutdown).
//     utils.StartServerWithGracefulShutdown(app)
// }


package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	AddUserRouter(v1)
	AddAuthRouter(v1)
	AddBookingRouter(v1)
}

package routes

import (
	"github.com/jailtonjunior94/bookings/api/infrastructure/ioc"
	"github.com/jailtonjunior94/bookings/api/presentation/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AddBookingRouter(router fiber.Router) {
	router.Get("/bookings", middlewares.Protected(), ioc.BookingController.Bookings)
	router.Post("/bookings", middlewares.Protected(), ioc.BookingController.CreateBooking)
}

package routes

import (
	"github.com/jailtonjunior94/bookings/api/infrastructure/ioc"

	"github.com/gofiber/fiber/v2"
)

func AddAuthRouter(router fiber.Router) {
	router.Post("/token", ioc.AuthController.Authenticate)
}