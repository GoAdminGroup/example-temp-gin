package router

import (
	"fmt"
	"github.com/GoAdminGroup/example-temp-gin/api"
	configProject "github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/example-temp-gin/router/interceptor"
	"github.com/GoAdminGroup/example-temp-gin/router/middleware"
	"github.com/GoAdminGroup/example-temp-gin/router/plugin"
	"github.com/GoAdminGroup/example-temp-gin/util/folder"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
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

	//config.ProbeViperString(
	//	"goAdmin.dataBases.default.driver",
	//	"goAdmin.dataBases.default.host",
	//	"goAdmin.dataBases.default.port",
	//	"goAdmin.dataBases.default.name",
	//	"goAdmin.dataBases.default.user",
	//	"goAdmin.dataBases.default.pwd",
	//	"goAdmin.dataBases.default.maxIdleCon",
	//	"goAdmin.dataBases.default.MaxOpenCon",
	//)
	monitorAPI(g)
	var env = viper.GetString("runmode")
	var routerDebug bool
	if env == "debug" || env == "test" {
		routerDebug = true
		checkPingServer(configProject.BaseURL())
	} else {
		routerDebug = false
	}

	var databaseList config.DatabaseList
	driver := viper.GetString("goAdmin.dataBases.default.driver")
	if driver == "sqlite" {
		dbFile := viper.GetString("goAdmin.dataBases.default.file")
		if dbFile == "" {
			return fmt.Errorf("in db drvier: %v , must set %v", driver, "goAdmin.dataBases.default.file")
		}
		exists, err := folder.PathExists(dbFile)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("in db drvier: %v, db path must exists at path: %v", driver, dbFile)
		}
		databaseList = config.DatabaseList{
			"default": {
				Driver:     driver,
				File:       dbFile,
				MaxIdleCon: viper.GetInt("goAdmin.dataBases.default.maxIdleCon"),
				MaxOpenCon: viper.GetInt("goAdmin.dataBases.default.MaxOpenCon"),
			},
		}
	} else {
		databaseList = config.DatabaseList{
			"default": {
				Driver:     driver,
				Host:       viper.GetString("goAdmin.dataBases.default.host"),
				Port:       viper.GetString("goAdmin.dataBases.default.port"),
				Name:       viper.GetString("goAdmin.dataBases.default.name"),
				User:       viper.GetString("goAdmin.dataBases.default.user"),
				Pwd:        viper.GetString("goAdmin.dataBases.default.pwd"),
				MaxIdleCon: viper.GetInt("goAdmin.dataBases.default.maxIdleCon"),
				MaxOpenCon: viper.GetInt("goAdmin.dataBases.default.MaxOpenCon"),
			},
		}
	}

	cfg := config.Config{
		Debug:     configProject.IsDebug(),
		Databases: databaseList,
		IndexUrl:  viper.GetString("goAdmin.indexUrl"),
		UrlPrefix: viper.GetString("goAdmin.urlPrefix"), // The url prefix of the website.
		// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
		Store: config.Store{
			Path:   viper.GetString("goAdmin.store.path"),
			Prefix: viper.GetString("goAdmin.store.prefix"),
		},
		Language:    configGoAdminLanguageByYaml(),
		ColorScheme: viper.GetString("goAdmin.color_scheme"),
		SqlLog:      viper.GetBool("goAdmin.sqlLog"),
	}

	// Plugin
	adminPlugin := plugin.AdminPlugin()
	examplePlugin := plugin.Example()

	if routerDebug {
		log.Infof("in envName %v ,will check server ping", env)
		delayDisplayViewRouterInfo()
	}

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

	static(g)

	if err := interceptor.Admin(eng, g); err != nil {
		return err
	}

	if err := api.PatchApi(eng, g); err != nil {
		return err
	}

	return nil
}
