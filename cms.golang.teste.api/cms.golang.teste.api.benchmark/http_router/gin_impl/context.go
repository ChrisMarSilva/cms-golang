package gin_impl

import (
	"github.com/ChrisMarSilva/cms.golang.teste.api.benchmark/http_router"

	"github.com/gin-gonic/gin"
)

func newContext(ctx *gin.Context) http_router.ContextRouter {
	return &ginContext{
		ctx: ctx,
	}
}

type ginContext struct {
	ctx *gin.Context
}

func (c *ginContext) JSON(code int, obj interface{}) error {
	c.ctx.JSON(code, obj)
	return nil
}
