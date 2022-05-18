package main

import (
	"log"

	"github.com/ChrisMarSilva/cms-golang-teste-api-fiber/route"
)

// docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=Hello123#" --name "sql1" -p 5401:1433 -v sql1data:/var/opt/mssql -d mcr.microsoft.com/mssql/server:2019-latest

// go mod init github.com/ChrisMarSilva/cms-golang-teste-api-fiber
// go get -u github.com/gofiber/fiber/v2
// go get -u github.com/gofiber/utils
// go get -u github.com/gofiber/fiber/middleware
// go get -u github.com/gofiber/jwt/v3
// go get -u github.com/golang-jwt/jwt/v4
// go get -u github.com/dgrijalva/jwt-go
// go get -u github.com/google/uuid
// go get -u gorm.io/gorm
// go get -u gorm.io/hints
// go get -u gorm.io/driver/mysql
// go get -u gorm.io/driver/sqlserver
// go get go.mongodb.org/mongo-driver/mongo
// go get github.com/uber/jaeger-client-go
// go get github.com/opentracing/opentracing-go
// go get github.com/jaegertracing/jaeger-client-go
// go get github.com/pkg/errors
// go get -u github.com/aschenmaker/fiber-opentracing
// go mod tidy

// go run main.go

func main() {

	sPort := "8001"
	log.Println("Rest API v2.0 - Fiber - Port ", sPort)

	app := route.NewRoutes()
	log.Fatal(app.Listen(":" + sPort))

}
