package router

import (
	"github.com/GoAdminGroup/example-temp-gin/config"
	"github.com/lexkong/log"
	"time"
)

func delayDisplayViewRouterInfo() {
	go delayViewRouterInfo()
}

func delayViewRouterInfo() {
	time.Sleep(time.Second * 5)
	log.Infof("=> delayDisplayViewRouterInfo Load view config BaseURL at: %v", config.BaseURL())
	log.Infof("login    at: %v%v", config.BaseURL(), "/admin ")
	//log.Infof("404 Page at: %v%v", config.BaseURL(), "/404.html")
}
