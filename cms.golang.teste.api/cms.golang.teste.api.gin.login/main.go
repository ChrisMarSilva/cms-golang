package main

import (
	"context" // "golang.org/x/net/context"
	"log"
	"net/http"
	"strings"

	"github.com/ChrisMarSilva/cms-golang-api-gin.login/models"
	// "github.com/ChrisMarSilva/cms-golang-api-gin.login/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// go mod init github.com/chrismarsilva/cms-golang-api-gin.login
// go mod tidy
// go get -u github.com/gin-gonic/gin
// go get -u githu.com/jackc/pgx/v4
// go get -u github.com/gofrs/uuid
// go get -u github.com/dgrijalva/jwt-go
// go get -u golang.org/x/crypto@v0.0.0-20210711020723-a769d52b0f97
// go run main.go

// exemplo: https://github.com/mikaelm1/offersapp

func main() {

	// conn, err := connectDB()
	// if err != nil {
	// 	return
	// }

	router := gin.Default()
	//router.Use(dbMiddleware(*conn))

	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"hello": "world"}) })
	//router.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	// routerGroup := router.Group("api/v1")
	// {
	// 	routerUserGroup := routerGroup.Group("users")
	// 	{
	// 		routerUserGroup.POST("register", routes.UsersRegister)
	// 		routerUserGroup.POST("login", routes.UsersLogin)
	// 	}
	// 	routerItemsGroup := routerGroup.Group("items")
	// 	{
	// 		routerItemsGroup.GET("index", routes.ItemsIndex)
	// 		routerItemsGroup.POST("create", authMiddleWare(), routes.ItemsCreate)
	// 		routerItemsGroup.PUT("update", authMiddleWare(), routes.ItemsUpdate)
	// 		routerItemsGroup.GET("sold_by_user", authMiddleWare(), routes.ItemsForSaleByCurrentUser)
	// 	}
	// }

	log.Println("Listem port 3000")
	router.Run(":3000")
}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:@localhost:5432/offersapp")
	if err != nil || conn == nil {
		log.Println("Error connecting to DB")
		log.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
