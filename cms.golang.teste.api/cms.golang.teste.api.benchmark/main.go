package main

import (
	router "github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router"
	chi "github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router/chi_impl"
	echo "github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router/echo_impl"
	fasthttp "github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router/fasthttp_routing_impl"
	fiber "github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router/fiber_impl"
	gin "github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router/gin_impl"
	mux "github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router/mux_impl"
	"github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/model"
	"net/http"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.benchmark
// go get -u github.com/gin-gonic/gin
// go get -u github.com/gofiber/fiber/v2
// go get -u github.com/valyala/fasthttp
// go get -u github.com/qiangxue/fasthttp-routing
// go get -u github.com/labstack/echo/v4
// go get -u github.com/go-chi/chi/v5
// go get -u github.com/gorilla/mux
// go mod tidy

// make up
// make run-fiber #used to run fiber benchmark
// make run-fasthttp # used to run fasthttp benchmark
// make run-gin # used to run gin benchmark
// make run-echo # used to run echo benchmark
// make run-mux # used to run mux benchmark
// make run-chi # used to run chi benchmark

// go run main.go

const portGin int = 8081
const portEcho int = 8082
const portMux int = 8083
const portFiber int = 8084
const portFastHttpRouting int = 8085
const portChi int = 8086

func main() {

	finished := make(chan bool)

	// ---------- INIT GIN ----------
	ginRouter := gin.New()
	ginRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "gin",
		})
	})
	go ginRouter.SERVE(portGin)

	// // ---------- INIT ECHO ----------
	echoRouter := echo.New()
	echoRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "echo",
		})
	})
	go echoRouter.SERVE(portEcho)

	// // ---------- INIT MUX ----------
	muxRouter := mux.New()
	muxRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "mux",
		})
	})
	go muxRouter.SERVE(portMux)

	// // ---------- INIT CHI ----------
	chiRouter := chi.New()
	chiRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "chi",
		})
	})
	go chiRouter.SERVE(portChi)

	// // ---------- INIT FIBER ----------
	fiberRouter := fiber.New()
	fiberRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "fiber",
		})
	})
	go fiberRouter.SERVE(portFiber)

	// // ---------- INIT FASTHTTP ROUTING ----------
	fasthttpRouter := fasthttp.New()
	fasthttpRouter.GET("/", func(c router.ContextRouter) error {
		return c.JSON(http.StatusOK, model.Router{
			Name: "fasthttp",
		})
	})
	go fasthttpRouter.SERVE(portFastHttpRouting)

	<-finished
}
