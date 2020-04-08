package api

import (
	"github.com/GoAdminGroup/example-temp-gin/api/demo"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
)

func PatchApi(eng *engine.Engine, g *gin.Engine) error {
	demo.PostData(g, eng)
	demo.StudentCount(g, eng)
	return nil
}
