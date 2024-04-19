package main

// go mod init github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth
// go get -u github.com/gin-gonic/gin
// go get -u github.com/stretchr/testify/assert
// go mod tidy
// go run main.go // go run .

// go install github.com/cosmtrek/air@latest
// air init
// air

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Server running at 0.0.0.0:8080 / localhost:8080") // 8080 // 3000

	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/", GetHome)

	main := r.Group("api/v1")
	{
		auth := main.Group("situacoes")
		{
			auth.GET("/:id", GetAuth)
			auth.POST("/", CreateAuth)
			auth.PUT("/", UpdateAuth)
			auth.DELETE("/:id", DeleteAuth)
		}
	}

	return r
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345") // Set example variable
		// before request
		c.Next()
		// after request
		log.Print(time.Since(t))       // latency := time.Since(t)
		log.Println(c.Writer.Status()) // access the status we are sending // status := c.Writer.Status()
	}
}

func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func CreateAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func UpdateAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func DeleteAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
