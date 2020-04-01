package router

import (
	"github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"time"
)

func delayDisplayViewRouterInfo() {
	go delayViewRouterInfo()
}

func delayViewRouterInfo() {
	time.Sleep(time.Second * 5)
	zlog.S().Infof("=> delayDisplayViewRouterInfo Load view config BaseURL at: %v", config.BaseURL())
	zlog.S().Infof("login    at: %v%v", config.BaseURL(), "/admin ")
}
