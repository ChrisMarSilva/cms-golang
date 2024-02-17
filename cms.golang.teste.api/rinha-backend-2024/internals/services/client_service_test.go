package services_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/dtos"
	"github.com/chrismarsilva/rinha-backend-2024/internals/handlers"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/chrismarsilva/rinha-backend-2024/internals/services"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/utils"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestUserrepoGetAll -v
// go test -run=XXX -bench . -benchmem

func GetRoutes() *fiber.App {
	viper.AddConfigPath("./")
	viper.SetConfigFile("../../cmd/api-server/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*", AllowCredentials: true}))
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	//Database
	driverDbWriter := databases.DatabasePostgres{}
	driverDbWriter.StartDbWriter()
	writer := driverDbWriter.GetDatabaseWriter()

	driverDbReader := databases.DatabasePostgres{}
	driverDbReader.StartDbReader()
	reader := driverDbReader.GetDatabaseReader()

	//Repository
	clientRepo := repositories.NewClientRepository(writer, reader)
	clientTransactionRepo := repositories.NewClientTransactionRepository(writer, reader)

	//Service
	clientServ := services.NewClientService(*clientRepo, *clientTransactionRepo)

	//Handle
	clientHandler := handlers.NewClientHandler(*clientServ)

	routes := app.Group("/clientes")
	routes.Post(":id/transacoes", clientHandler.CreateTransaction)
	routes.Get("/:id/extrato", clientHandler.GetExtract)
	app.Use(NotFound)

	return app
}

func NotFound(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}

func TestRouteTransacao(t *testing.T) {
	app := GetRoutes()

	payload := dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "Teste"}
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(payload)

	req := httptest.NewRequest("POST", "/clientes/1/extrato", &body)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON) // MIMEApplicationJSONCharsetUTF8

	resp, err := app.Test(req)
	if 200 != resp.StatusCode {
		body, _ := ioutil.ReadAll(resp.Body)
		println(string(body))
	}

	utils.AssertEqual(t, nil, err, "TransacaoErr")
	utils.AssertEqual(t, 200, resp.StatusCode, "TransacaoStatusCode")
}

func TestRouteExtrato(t *testing.T) {
	app := GetRoutes()

	req := httptest.NewRequest("GET", "/clientes/1/extrato", nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON) // MIMEApplicationJSONCharsetUTF8

	resp, err := app.Test(req)
	if 200 != resp.StatusCode {
		body, _ := ioutil.ReadAll(resp.Body)
		println(string(body))
	}

	utils.AssertEqual(t, nil, err, "ExtratoErr")
	utils.AssertEqual(t, 200, resp.StatusCode, "ExtratoStatusCode")
}

func TestRoutes(t *testing.T) {
	app := GetRoutes()

	payload := dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "Teste"}
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(payload)

	tests := []struct {
		description  string // description of the test case
		method       string // type path to test
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		body         io.Reader
	}{
		{description: "Transações status 200 Ok", method: "GET", route: "/clientes/1/transacoes", expectedCode: 200, body: &body},
		{description: "Transações status 404 NotFound", method: "GET", route: "/clientes/6/transacoes", expectedCode: 404, body: &body},
		{description: "Extrato status 200 Ok", method: "GET", route: "/clientes/1/extrato", expectedCode: 200, body: nil},
		{description: "Extrato status 404 NotFound", method: "GET", route: "/clientes/6/extrato", expectedCode: 404, body: nil},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)

		if test.expectedCode != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			// utils.AssertEqual(t, nil, err, test.description+" app.Test(req)")
			// utils.AssertEqual(t, ":param", body, test.description+" body")
			println(string(body))
		} // if test.expectedCode != resp.StatusCode {

		utils.AssertEqual(t, nil, err, test.description+" Error")
		utils.AssertEqual(t, test.expectedCode, resp.StatusCode, test.description+" StatusCode")
	}
}

func BenchmarkRouteTransacoes(b *testing.B) {
	app := GetRoutes()

	payload := dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "Teste"}
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(payload)

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/clientes/1/transacoes", &body)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := app.Test(req)
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}

		utils.AssertEqual(b, nil, err, "TransacaoErr")
		utils.AssertEqual(b, 200, resp.StatusCode, "TransacaoStatusCode")
	}
}

func BenchmarkRouteExtrato(b *testing.B) {
	app := GetRoutes()

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/clientes/1/extrato", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		resp, err := app.Test(req)
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}

		utils.AssertEqual(b, nil, err, "ExtratoErr")
		utils.AssertEqual(b, 200, resp.StatusCode, "ExtratoStatusCode")
	}
}
