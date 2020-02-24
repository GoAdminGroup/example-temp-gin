package demo

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostData(g *gin.Engine, eng *engine.Engine) gin.IRoutes {
	return g.POST("/admin/postData", func(ctx *gin.Context) {

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

		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": "<h2>hello world</h2>",
		})
	})
}
