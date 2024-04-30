package main

// import (
// 	"time"

// 	"github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth/internals/server"
// 	"github.com/gofiber/fiber/v2/middleware/session"
// 	"github.com/gofiber/storage/sqlite3"
// )

// var (
// 	storage = sqlite3.New(sqlite3.Config{Database: "./banco.db"})

// 	store = session.New(session.Config{
// 		Expiration: 24 * time.Hour,
// 		KeyLookup:  "cookie:session_id",
// 		Storage:    storage,
// 	})
// )

// func init() {
// 	// config, err := initializers.LoadConfig(".")
// 	// if err != nil {
// 	// 	log.Fatalln("Failed to load environment variables! \n", err.Error())
// 	// }
// 	// initializers.ConnectDB(&config)
// }

// func main() {
// 	app := server.NewServer()
// 	app.Initialize()
// }

/*


func loadDatabase() {
    database.InitDb()
    database.Db.AutoMigrate(&model.Role{})
    database.Db.AutoMigrate(&model.User{})
    seedData()
}

// load seed data into the database
func seedData() {
    var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
    var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
    database.Db.Save(&roles)
    database.Db.Save(&user)
}


func serveApplication() {
    router := gin.Default()
    authRoutes := router.Group("/auth/user")
    authRoutes.POST("/register", controller.Register)
    authRoutes.POST("/login", controller.Login)

    router.Run(":8000")
    fmt.Println("Server running on port 8000")
}


func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
	database.InitDb()
}

func serveApplication() {
	router := gin.Default()

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}



    models.InitDB(config)
    routes.AuthRoutes(r)
    r.Run(":8080")



 http://localhost:1323/swagger/index.html,


func main() {
    logger, _ := zap.NewProduction()
    ctx := context.Background()
    s := NewServiceA(logger)
    ctx = zax.Set(ctx, logger, []zap.Field{zap.String("trace_id", "my-trace-id")})
    s.funcA(ctx)
}

type ServiceA struct {
logger *zap.Logger
}

func NewServiceA(logger *zap.Logger) *ServiceA {
    return &ServiceA{
        logger: logger,
    }
}

func (s *ServiceA) funcA(ctx context.Context) {
    s.logger.Info("func A") // it does not contain trace_id, you need to add it manually
    zax.Get(ctx).Info("func A") // it will logged with "trace_id" = "my-trace-id"
}


// edit main.go
func serveApplication() {
        adminRoutes := router.Group("/admin")
    adminRoutes.Use(util.JWTAuth())
    adminRoutes.GET("/users", controller.GetUsers)
    adminRoutes.GET("/user/:id", controller.GetUser)
    adminRoutes.PUT("/user/:id", controller.UpdateUser)
    adminRoutes.POST("/user/role", controller.CreateRole)
    adminRoutes.GET("/user/roles", controller.GetRoles)
    adminRoutes.PUT("/user/role/:id", controller.UpdateRole)
}

    db, err := run()
    if (err != nil){
        log.Fatal(err)
    }
    defer db.SQL.Close()
go install github.com/cortesi/modd/cmd/modd@latest
modd

*/
