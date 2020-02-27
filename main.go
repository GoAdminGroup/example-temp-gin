package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/GoAdminGroup/example-temp-gin/router"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // adapter
	_ "github.com/GoAdminGroup/themes/adminlte"      // theme

	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sqlite driver
	//_ "github.com/go-sql-driver/mysql" // mysql driver
	//_ "github.com/lib/pq" // postgresql driver

	configProject "github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "api server config file path.")
)

func main() {

	pflag.Parse()

	// init config
	if err := configProject.Init(*cfg); err != nil {
		fmt.Printf("Error, run service not use -c or config yaml error, more info: %v\n", err)
		panic(err)
	}
	fmt.Printf("%s -> %v at time: %v\n", "start service", viper.GetString("name"), time.Now().String())

	// Set gin mode.
	runMode := viper.GetString("runmode")
	gin.SetMode(runMode)

	g := gin.Default()

	gin.DefaultWriter = ioutil.Discard

	eng := engine.Default()

	var middlewareList []gin.HandlerFunc

	if err := router.Load(
		// Cores.
		eng, g,
		// middlewareList.
		middlewareList...,
	); err != nil {
		panic(err)
	}

	log.Infof("Start to listening the incoming requests on http address: %v", viper.GetString("addr"))
	log.Infof("Sever name: %v , has start!", viper.GetString("name"))

	if configProject.IsDebug() {
		err := http.ListenAndServe(viper.GetString("addr"), g)
		if err != nil {
			log.Errorf(err, "server run error %v", err)
		} else {
			log.Infof("server run success!")
		}
	} else {
		err := endless.ListenAndServe(viper.GetString("addr"), g)
		if err != nil {
			log.Errorf(err, "server run error %v", err)
		} else {
			log.Infof("server run success!")
		}
	}
}
