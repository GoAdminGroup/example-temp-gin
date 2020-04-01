package router

import (
	"github.com/GoAdminGroup/example-temp-gin/api"
	"github.com/GoAdminGroup/example-temp-gin/components/loginComp"
	configProject "github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/GoAdminGroup/example-temp-gin/router/interceptor"
	"github.com/GoAdminGroup/example-temp-gin/router/middleware"
	"github.com/GoAdminGroup/example-temp-gin/router/plugin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Load(eng *engine.Engine, g *gin.Engine, mw ...gin.HandlerFunc) error {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.LoggerMiddleware())
	g.Use(mw...)
	//// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		noRouterBiz(c)
	})

	monitorAPI(g)

	global := configProject.Global()

	cfg := goAdminConfig(global)

	if global.GoAdmin.UseCustom {
		goAdminCustom(&cfg, global.GoAdmin.Custom)
	}

	if global.GoAdmin.Captcha {
		template.AddLoginComp(loginComp.GetLoginComponent())
	}

	// Plugin
	adminPlugin := plugin.AdminPlugin()
	examplePlugin := plugin.Example()

	err := eng.AddConfig(cfg).
		AddPlugins(adminPlugin, examplePlugin).
		Use(g)
	if err != nil {
		return err
	}

	err = initDBConnection(eng)
	if err != nil {
		return err
	}

	err = static(g)
	if err != nil {
		return err
	}

	if err := interceptor.Admin(eng, g); err != nil {
		return err
	}

	if err := api.PatchApi(eng, g); err != nil {
		return err
	}

	if configProject.IsDebug() {
		zlog.S().Infof("in envName %v ,will check server ping", viper.GetString("run_mode"))
		delayDisplayViewRouterInfo()
	}
	checkPingServer(configProject.BaseURL())
	return nil
}
