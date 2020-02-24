package index

import "github.com/spf13/viper"

var cfgBase *baseConfig

type baseConfig struct {
	Title string
}

func ConfigBase() *baseConfig {
	if cfgBase == nil {
		cfgBase = &baseConfig{
			Title: viper.GetString("goAdmin.dashBoard.title"),
		}
	}
	return cfgBase
}
