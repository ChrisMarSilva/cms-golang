package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"

	//"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	// "github.com/rs/zerolog"
	// sqldblogger "github.com/simukti/sqldb-logger"
	// "github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

// go mod init github.com/chrismarsilva/cms-golang-api-fiber.basic
// go get -u github.com/gofiber/fiber/v2
// go get -u github.com/gofiber/utils
// go get -u github.com/gofiber/fiber/middleware
// go get -u github.com/go-sql-driver/mysql
// go get -u github.com/denisenkom/go-mssqldb
// go get -u -v github.com/simukti/sqldb-logger
// go get -u github.com/simukti/sqldb-logger/logadapter/zerologadapter
// go get -u github.com/arsmn/fiber-swagger/v2
// go mod tidy

// https://github.com/arsmn/fiber-swagger
// swag init
// swag fmt

// go run main.go

func main() {

	// //driverName := "mysql"
	// //dataSourceName  := "root:senha@tcp(localhost:3306)/database"

	// driverName := "mssql"
	// dataSourceName := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_TNB"

	// db, err := sql.Open(driverName, dataSourceName)
	// if err != nil {
	// 	log.Fatalln("sql.Open", err)
	// }
	// defer db.Close() // adiar o fechamento até depois que a função principal terminar

	// db.SetMaxIdleConns(10000) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	// db.SetMaxOpenConns(10000) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	// db.SetConnMaxIdleTime(time.Hour)
	// db.SetConnMaxLifetime(time.Minute * 30) // 24 *time.Hour // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	// db.Exec("TRUNCATE TABLE TBTESTE_API")

	app := fiber.New()
	//app.Use(logger.New(logger.Config{Format: "[${time}]: ${ip} ${status} ${latency} ${method} ${path} Error: ${error}\n"}))
	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	//app.Get("/swagger/*any", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok fiber 2")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok ping")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok health")
	})

	app.Get("/data", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok data")
	})

	app.Get("/getall", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("ok getall")
	})

	// custom
	// app.Get("/swagger/*", swagger.New(swagger.Config{
	// 	URL:               "http://example.com/doc.json",
	// 	DeepLinking:       false,
	// 	DocExpansion:      "none",
	// 	OAuth:             &swagger.OAuthConfig{AppName: "OAuth Provider", ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2"},
	// 	OAuth2RedirectUrl: "http://localhost:8002/swagger/oauth2-redirect.html",
	// }))

	//handler := NewHandler(db)
	//app.Get("/ok", handler.Teste)
	//app.Get("/delete", handler.Delete)
	//app.Get("/insert/:id", handler.Insert)

	//SetupApiV1(app)

	log.Fatal(app.Listen(":8002"))

}

func SetupApiV1(app *fiber.App) {
	v1 := app.Group("/v1")
	SetupTodosRoutes(v1)
}

func SetupTodosRoutes(grp fiber.Router) {
	todosRoutes := grp.Group("/todos")
	todosRoutes.Get("/", GetTodos)
	todosRoutes.Get("/:id", GetTodo)
	todosRoutes.Post("/", CreateTodo)
	todosRoutes.Patch("/:id", UpdateTodo)
	todosRoutes.Delete("/:id", DeleteTodo)
}

func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(todos)
}

func GetTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse id"})
	}

	for _, todo := range todos {
		if todo.Id == id {
			return ctx.Status(fiber.StatusOK).JSON(todo)
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON("")
}

func CreateTodo(ctx *fiber.Ctx) error {

	type request struct {
		Name string `json:"name"`
	}

	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
	}

	todo := &Todo{Id: len(todos) + 1, Name: body.Name, Completed: false}
	todos = append(todos, todo)

	return ctx.Status(fiber.StatusCreated).JSON(todo)
}

func UpdateTodo(ctx *fiber.Ctx) error {

	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse id"})
	}

	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse body"})
	}

	var todo *Todo
	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo == nil {
		return ctx.Status(fiber.StatusNotFound).JSON("")
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func DeleteTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")

	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse id"})
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:i], todos[i+1:]...)
			return ctx.Status(fiber.StatusNoContent).JSON("")
		}
	}

	return ctx.Status(fiber.StatusNotFound).JSON("")
}

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Name: "Walk the dog", Completed: false},
	{Id: 2, Name: "Walk the cat", Completed: false},
}

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Teste(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("ok fiber 2")
}

func (h *Handler) Delete(c *fiber.Ctx) error {

	h.db.Exec("TRUNCATE TABLE TBTESTE_API")

	return c.Status(fiber.StatusOK).JSON("")
}

func (h *Handler) Insert(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		log.Println("id empty")
		return c.Status(fiber.StatusBadRequest).JSON("id empty")
	}

	for i := 0; i < 10000; i++ {
		err := h.db.Ping()
		if err != nil {
			//log.Println("sql.Ping.erro", err, "id:", id, "i:", i)
			//return c.Status(fiber.StatusBadRequest).JSON("sql.Exec:" + err.Error())
			time.Sleep(time.Millisecond * 100)
			continue
		}
		// if i > 0 {
		// 	log.Println("sql.Ping.ok", "id:", id, "i:", i)
		// }
		break
	}

	for i := 0; i < 10; i++ {

		t := time.Now()

		result, err := h.db.Exec("INSERT INTO TBTESTE_API (NOME, DTHR, SITUACAO) VALUES (?,?,?)", "Teste "+id, t.Format("20060102150405"), "A")
		if err != nil {
			// log.Println("sql.Exec", err)
			// return c.Status(fiber.StatusBadRequest).JSON("sql.Exec:" + err.Error())
			time.Sleep(time.Millisecond * 100)
			continue
		}

		if rowsAffected, _ := result.RowsAffected(); rowsAffected != 1 {
			// log.Println("sql.RowsAffected", err)
			//return c.Status(fiber.StatusBadRequest).JSON("sql.RowsAffected:" + err.Error())
			time.Sleep(time.Millisecond * 100)
			continue
		}

		return c.Status(fiber.StatusOK).JSON("")

	}

	return c.Status(fiber.StatusBadRequest).JSON("ssss")
}
