package service_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	database "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/databases"
	"github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories/user"
	service "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestUserrepoGetAll -v
// go test -run=XXX -bench . -benchmem

func TestRoute(t *testing.T) {

	type UserLocal struct {
		Nome string `json:"nome"`
	}
	userLocal := UserLocal{Nome: "Pessoa Teste"}
	var userCreate bytes.Buffer
	json.NewEncoder(&userCreate).Encode(userLocal)
	var userUpdate bytes.Buffer
	json.NewEncoder(&userUpdate).Encode(userLocal)

	//println(userCreate)
	//body1, _ := ioutil.ReadAll(&userCreate)
	//println(string(body1))

	tests := []struct {
		description  string // description of the test case
		method       string // type path to test
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		body         io.Reader
	}{
		{description: "get HTTP status 200", method: "GET", route: "/hello", expectedCode: 200, body: nil},
		{description: "get HTTP status 404, when route is not exists", method: "GET", route: "/not-found", expectedCode: 404, body: nil},
		{description: "GetAll HTTP status 404", method: "GET", route: "/api/v1/user", expectedCode: 200, body: nil},
		{description: "Get HTTP status 404", method: "GET", route: "/api/v1/user/A6D53851-13BA-4E2E-8FFA-00138D02A281", expectedCode: 200, body: nil},
		{description: "Create HTTP status 404", method: "POST", route: "/api/v1/user", expectedCode: 200, body: &userCreate}, // strings.NewReader("{'nome': 'Pessoa Teste'}")
		{description: "Update HTTP status 404", method: "PUT", route: "/api/v1/user/A6D53851-13BA-4E2E-8FFA-00138D02A281", expectedCode: 200, body: &userUpdate},
		{description: "Delete HTTP status 404", method: "DELETE", route: "/api/v1/user/ACCEA3FE-373C-4F6C-A317-7DC14C547D90", expectedCode: 200, body: nil},
	}

	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	driverDB := database.DatabaseSQLServer{} // DatabaseMySQL // DatabaseSQLServer
	driverDB.StartDB()
	db := driverDB.GetDatabase()

	userRepo := user.NewUserRepositoryMSSQL(nil, db)
	userService := service.NewUserService(*userRepo)

	api := app.Group("/api")                          // /api
	v1 := api.Group("/v1")                            // /api/v1
	v1.Get("/user", userService.HandlerGetAll)        // /api/v1/user
	v1.Get("/user/:id", userService.HandlerGet)       // /api/v1/user
	v1.Post("/user", userService.HandlerCreate)       // /api/v1/user
	v1.Put("/user/:id", userService.HandlerUpdate)    // /api/v1/user
	v1.Delete("/user/:id", userService.HandlerDelete) // /api/v1/user

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req) // (req, 1)
		// resp, err := app.Test(httptest.NewRequest("GET", test.route, nil))

		if test.expectedCode != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			// utils.AssertEqual(t, nil, err, "app.Test(req)")
			// utils.AssertEqual(t, ":param", body)
			println(string(body))
		} // if test.expectedCode != resp.StatusCode {

		utils.AssertEqual(t, nil, err, test.description+" Error")
		utils.AssertEqual(t, test.expectedCode, resp.StatusCode, test.description)
		//assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		//println(test.expectedCode, resp.StatusCode, test.description)

	} // for _, test := range tests {
}

func BenchmarkRepoGetAll(b *testing.B) {
	driverDB := database.DatabaseSQLServer{} // DatabaseMySQL // DatabaseSQLServer
	driverDB.StartDB()
	db := driverDB.GetDatabase()
	userRepo := user.NewUserRepositoryMSSQL(nil, db)
	userService := service.NewUserService(*userRepo)
	app := fiber.New()
	api := app.Group("/api")                   // /api
	v1 := api.Group("/v1")                     // /api/v1
	v1.Get("/user", userService.HandlerGetAll) // /api/v1/user
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/api/v1/user", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}
		utils.AssertEqual(b, nil, err, "GetAllError")
		utils.AssertEqual(b, 200, resp.StatusCode, "GetAll")
	}
}

func BenchmarkRepoGet(b *testing.B) {
	driverDB := database.DatabaseSQLServer{} // DatabaseMySQL // DatabaseSQLServer
	driverDB.StartDB()
	db := driverDB.GetDatabase()
	userRepo := user.NewUserRepositoryMSSQL(nil, db)
	userService := service.NewUserService(*userRepo)
	app := fiber.New()
	api := app.Group("/api")                    // /api
	v1 := api.Group("/v1")                      // /api/v1
	v1.Get("/user/:id", userService.HandlerGet) // /api/v1/user
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/api/v1/user/A6D53851-13BA-4E2E-8FFA-00138D02A281", nil)
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}
		utils.AssertEqual(b, nil, err, "GetError")
		utils.AssertEqual(b, 200, resp.StatusCode, "Get")
	}
}
