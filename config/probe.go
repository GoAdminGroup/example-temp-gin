package config

import (
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/spf13/viper"
)

func ProbeViperString(configs ...string) {
	if len(configs) > 0 {
		for _, config := range configs {
			zlog.S().Debugf("probe %v -> %v", config, viper.GetString(config))
		}
	}
}
