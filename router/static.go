package router

import (
	"github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/example-temp-gin/util/tools"
	"github.com/gin-gonic/gin"
	"os"
)

func static(g *gin.Engine) error {
	// set goAdmin store
	storePath := config.Global().GoAdmin.Store.Path
	g.Static(config.Global().GoAdmin.Store.Prefix, storePath)

	exists, err := tools.PathExists(storePath)
	if err != nil {
		return err
	}
	if !exists {
		err := os.Mkdir(storePath, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}
