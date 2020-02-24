package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func static(g *gin.Engine) {
	// set goAdmin store
	g.Static(viper.GetString("goAdmin.store.prefix"), viper.GetString("goAdmin.store.path"))
}
