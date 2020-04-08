package demo

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
)

func demo(g *gin.Engine, eng *engine.Engine) gin.IRoutes {
	bizPath := "/demo"
	return g.GET(bizPath, func(ctx *gin.Context) {

	})
}
