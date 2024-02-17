// package routes

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"cms.golang.tnb.api/controllers"
// 	"github.com/gin-contrib/cache"
// 	"github.com/gin-contrib/cache/persistence"
// 	"github.com/gin-gonic/gin"
// )

// func ConfigRoutes(router *gin.Engine) *gin.Engine {

// 	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, "pong5") })
// 	router.GET("/ping1", func(c *gin.Context) { c.String(200, "pong "+fmt.Sprint(time.Now().Unix())) })

// 	store := persistence.NewInMemoryStore(time.Second)
// 	router.GET("/ping2", cache.CachePage(store, time.Minute, func(c *gin.Context) { c.String(200, "pong "+fmt.Sprint(time.Now().Unix())) }))

// 	main := router.Group("api/v1")
// 	{
// 		situacoes := main.Group("situacoes")
// 		{
// 			situacaoRepo := controllers.New()
// 			situacoes.GET("/", situacaoRepo.GetSituacoes)
// 			situacoes.GET("/:id", situacaoRepo.GetSituacao)
// 			// situacoes.POST("/", situacaoRepo.CreateUser)
// 			// situacoes.PUT("/", situacaoRepo.UpdateUser)
// 			// situacoes.DELETE("/:id", situacaoRepo.DeleteUser)
// 		}
// 	}
// 	return router
// }

// // r.GET("/situacoes", situacaoRepo.GetSituacoes)
// // r.GET("/situacoes/:codigo", situacaoRepo.GetSituacao)
// // // r.POST("/situacoes", situacaoRepo.CreateUser)
// // // r.PUT("/situacoes/:codigo", situacaoRepo.UpdateUser)
// // // r.DELETE("/situacoes/:codigo", situacaoRepo.DeleteUser)
