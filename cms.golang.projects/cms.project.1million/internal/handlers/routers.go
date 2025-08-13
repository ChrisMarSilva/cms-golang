package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Config        *utils.Config
	PersonHandler *PersonHandler
}

func NewRouter(config *utils.Config, personHandler *PersonHandler) *Router {
	return &Router{
		Config:        config,
		PersonHandler: personHandler,
	}
}

func (r Router) Listen() error {
	gin.SetMode(r.Config.GinMode)

	router := gin.Default()
	router.Use(r.ErrorsMiddleware())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.BestSpeed))

	router.GET("/health", r.HealthHandler)

	v1Group := router.Group("/v1")
	{
		personGroup := v1Group.Group("/person")
		{
			personGroup.POST("", r.PersonHandler.Add)
			personGroup.GET("/", r.PersonHandler.GetAll)
			personGroup.GET("/:id", r.PersonHandler.GetByID)
			personGroup.GET("/count", r.PersonHandler.GetCount)
		}
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", r.Config.UriPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s.ListenAndServe()
}

func (r Router) ErrorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		}
	}
}

func (r Router) HealthHandler(c *gin.Context) {
	//c.JSON(http.StatusOK, "ok")
	c.JSON(http.StatusOK, r.Config.NameApi)
}
