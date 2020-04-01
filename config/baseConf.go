package config

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/GoAdminGroup/example-temp-gin/util/sys"
	"github.com/spf13/viper"
)

var baseConf *baseConfig

type baseConfig struct {
	IsDebug    bool
	EnvName    string
	BaseURL    string
	BaseHost   string
	ApiVersion string
	SSLEnable  bool
}

func BaseConfig() *baseConfig {
	return baseConf
}

func EnvName() string {
	return baseConf.EnvName
}

func BaseURL() string {
	return baseConf.BaseURL
}

func BaseHost() string {
	return baseConf.BaseHost
}

func IsDebug() bool {
	return baseConf.IsDebug
}

// read default config by conf/config.yaml
// can change by CLI by `-c`
// this config can config by ENV
//	ENV_WEB_HTTPS_ENABLE=false
//	ENV_AUTO_HOST=true
//	ENV_WEB_HOST 127.0.0.1:8000
func initBaseConf() {
	var env = viper.GetString("run_mode")
	var confDebug bool
	if env == "debug" || env == "test" {
		confDebug = true
	} else {
		confDebug = false
	}

	ssLEnable := false
	if viper.GetBool(defaultEnvHttpsEnable) {
		ssLEnable = true
	} else {
		ssLEnable = viper.GetBool("sslEnable")
	}
	apiBase := viper.GetString("api_url")

	uri, err := url.Parse(apiBase)
	if err != nil {
		panic(err)
	}
	var apiHost string
	apiHost = uri.Host
	baseHOSTByEnv := viper.GetString(defaultEnvHost)
	if baseHOSTByEnv != "" {
		uri.Host = baseHOSTByEnv
		apiHost = baseHOSTByEnv
		apiBase = uri.String()
	} else {
		isAutoHost := viper.GetBool(defaultEnvAutoGetHost)
		zlog.S().Debugf("isAutoHost %v", isAutoHost)
		if isAutoHost {
			ipv4, err := sys.NetworkLocalIP()
			if err == nil {
				addrStr := viper.GetString("addr")
				var proc string
				if ssLEnable {
					proc = "https"
				} else {
					proc = "http"
				}
				apiBase = fmt.Sprintf("%v://%v%v", proc, ipv4, addrStr)
				apiHost = fmt.Sprintf("%v%v", ipv4, addrStr)
			}
		}
	}

	if ssLEnable {
		apiBase = strings.Replace(apiBase, "http://", "https://", 1)
	}
	zlog.S().Debugf("config.BaseConfig()\nIsDebug: %v\nEnvName: %v\nBaseURL: %v\nBaseHost: %v\nApiVersion: %v\nSSLEnable: %v",
		confDebug, env, apiBase, apiHost, viper.GetString("api_version"), ssLEnable)
	baseConf = &baseConfig{
		IsDebug:    confDebug,
		EnvName:    env,
		BaseURL:    apiBase,
		BaseHost:   apiHost,
		ApiVersion: viper.GetString("api_version"),
		SSLEnable:  ssLEnable,
	}
}
