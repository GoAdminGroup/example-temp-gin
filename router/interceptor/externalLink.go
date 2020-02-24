package interceptor

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func externalLink(g *gin.Engine, eng *engine.Engine) gin.IRoutes {
	return g.GET("/admin/external_link", func(ctx *gin.Context) {
		user, ok := eng.User(ctx)

		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "auth fail",
			})
			return
		}

		if !user.CheckPermission("*") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "没有权限",
			})
			return
		}
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return types.Panel{
				Content: `<iframe style="width:100%;height:800px;" src="https://gitee.com/go-admin/go-admin"></iframe>`,
			}, nil
		})
	})
}
