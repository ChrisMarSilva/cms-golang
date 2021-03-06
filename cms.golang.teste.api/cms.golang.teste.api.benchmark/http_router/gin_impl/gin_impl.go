package gin_impl

import (
	"fmt"
	"github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ginImpl struct {
	dispatch *gin.Engine
}

func New() http_router.Router {
	gin.SetMode(gin.ReleaseMode)
	return &ginImpl{dispatch: gin.New()}
}

func (g *ginImpl) GET(uri string, resolver http_router.ResolveHandler) {
	g.dispatch.GET(uri, func(c *gin.Context) {
		ctx := newContext(c)
		_ = resolver(ctx)
	})
}

func (g *ginImpl) SERVE(port int) {
	portString := strconv.Itoa(port)
	fmt.Println("serving with GIN on ", port)
	log.Fatal(g.dispatch.Run(":" + portString))
}
