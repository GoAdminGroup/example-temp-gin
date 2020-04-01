package config

import (
	"github.com/GoAdminGroup/example-temp-gin/pkg/zlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Monitor configuration changes and hot loaders
func (c *ConfFile) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		zlog.S().Debugf("config file changed: %s", e.Name)
		err := viper.Unmarshal(&global)
		if err != nil {
			zlog.S().Errorf("viper.Unmarshal error: %v", err)
		}
	})
}
