package config

import (
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func ProbeViperString(configs ...string) {
	if len(configs) > 0 {
		for _, config := range configs {
			log.Debugf("probe %v -> %v", config, viper.GetString(config))
		}
	}
}
