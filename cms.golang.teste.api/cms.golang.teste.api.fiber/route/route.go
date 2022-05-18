package route

import (
	"crypto/rsa"
	"time"

	database "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/databases"
	handler "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/handlers"
	repository "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories"
	"github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories/user"
	service "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"gorm.io/gorm"

	//"github.com/dgrijalva/jwt-go"
	//"github.com/gofiber/fiber/v2/middleware/csrf"
	jwtware "github.com/gofiber/jwt/v3"

	"github.com/aschenmaker/fiber-opentracing/fjaeger"
	jconfig "github.com/uber/jaeger-client-go/config"
)

const jwtSecret = "asecret#crss"

var privateKey *rsa.PrivateKey

func NewRoutes() *fiber.App {

	driverDB := database.DatabaseSQLServer{} // DatabaseSQLServer // DatabaseMongo
	driverDB.StartDB()
	db := driverDB.GetDatabase()

	logRepo := repository.NewLogMonitorRepositoryMSSQL("API", "FBR", db)
	//logRepo.Inserir("NewRoutes", "1", "Inicializado com Sucesso")

	userRepo := user.NewUserRepositoryMSSQL(nil, db, *logRepo) // NewUserRepositoryMSSQL(nil, db, *logRepo) // NewUserRepositoryMongo(driverDB.Context, db, *logRepo)
	userServ := service.NewUserService(*userRepo, *logRepo)
	userHandler := handler.NewUserHandler(*userServ, *logRepo)

	defaultHandler := handler.NewDefaultHandler()

	app := fiber.New()

	// SetupConfigRoutes(app)

	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
	// 		code := fiber.StatusInternalServerError // Status code defaults to 500
	// 		var msg string
	// 		if e, ok := err.(*fiber.Error); ok { // Retrieve the custom status code if it's an fiber.*Error
	// 			code = e.Code
	// 			msg = e.Message
	// 		}
	// 		if msg == "" {
	// 			msg = "cannot process the http call"
	// 		}
	// 		err = ctx.Status(code).JSON(internalError{ Message: msg }) // Send custom error page
	// 		return nil
	// 	},
	// })

	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	//app.Use(cors.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*", AllowCredentials: true}))
	//app.Use(csrf.New())
	app.Use(encryptcookie.New(encryptcookie.Config{Key: "VWY5XBVap84Zpd0ckbT1reTl0NM6pz7R"}))
	//app.Use(logger.New())
	app.Use(logger.New(logger.Config{Format: "[${time}]: ${ip} ${status} ${latency} ${method} ${path} Error: ${error}\n"}))
	//app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n", TimeFormat: "02-Jan-2006", TimeZone:   "Asia/Jakarta" }))
	app.Use(pprof.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	cfg := fjaeger.Config{
		ServiceName:      "cms.golang.api.fiber",
		Sampler:          &jconfig.SamplerConfig{Type: "const", Param: 1},
		Reporter:         &jconfig.ReporterConfig{LogSpans: true, BufferFlushInterval: 1 * time.Second},
		EnableRPCMetrics: true,
		PanicOnError:     false,
	}

	fjaeger.New(cfg)

	// app.Use(fibertracing.New(fibertracing.Config{
	// 	Tracer: opentracing.GlobalTracer(),
	// 	OperationName: func(ctx *fiber.Ctx) string {
	// 		return "TEST:  HTTP " + ctx.Method() + " URL: " + ctx.Path()
	// 	},
	// }))

	//SetupUsersRoutes(app, db, logRepo)
	app.Get("/api/v1/user", userHandler.GetAll)        // /api/v1/user
	app.Get("/api/v1/user/:id", userHandler.Get)       // /api/v1/user
	app.Post("/api/v1/user", userHandler.Create)       // /api/v1/user
	app.Put("/api/v1/user/:id", userHandler.Update)    // /api/v1/user
	app.Delete("/api/v1/user/:id", userHandler.Delete) // /api/v1/user

	//SetupDefaultRoutes(app)
	app.Get("/", timeout.New(defaultHandler.Index, 5*time.Second))
	app.Get("/teste", defaultHandler.Teste)
	app.Get("/dashboard", monitor.New())
	app.Get("/login", defaultHandler.Login)   // /login
	app.Get("/logout", defaultHandler.Logout) // /logout
	app.Get("/restricted1", defaultHandler.Restricted1)
	app.Get("/restricted2", AuthRequired, defaultHandler.Restricted2)
	public := app.Group("/public")
	public.Get("/", defaultHandler.Public)
	private := app.Group("/private")
	private.Use(HandlerAuthRequired())
	//private.Use(jwtware.New(jwtware.Config{SigningKey: []byte(jwtSecret)}))
	private.Get("/", defaultHandler.Private)
	app.Use(defaultHandler.NotFound)
	public.Use(defaultHandler.NotFound)
	private.Use(defaultHandler.NotFound)

	return app
}

func SetupConfigRoutes(app *fiber.App) {

}

func SetupUsersRoutes(app *fiber.App, db *gorm.DB, logRepo *repository.LogMonitorRepositoryMSSQL) {

}

func SetupDefaultRoutes(app *fiber.App) {

}

func HandlerAuthRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		TokenLookup:    "header:Authorization",
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		// ErrorHandler: func(c *fiber.Ctx, err error) {
		// 	err = c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		// },
		SigningKey:    []byte(jwtSecret),
		SigningMethod: "HS256", // jwt.SigningMethodHS256.Name
	})
}

func AuthRequired(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		TokenLookup:    "header:Authorization",
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized", "msg": err.Error()})
		// },
		SigningKey:    []byte(jwtSecret),
		SigningMethod: "HS256", // jwt.SigningMethodHS256.Name
	})(ctx)
}

func AuthError(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized", "msg": err.Error()})
	return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}
