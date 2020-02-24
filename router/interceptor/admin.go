package interceptor

import (
	"github.com/GoAdminGroup/example-temp-gin/pages/index"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

func Admin(eng *engine.Engine, g *gin.Engine) error {

	g.GET("/admin", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return index.DashboardPage()
		})
	})

	externalLink(g, eng)

	return nil
}
