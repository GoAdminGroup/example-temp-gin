package router

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/spf13/viper"
)

// default use language.CN
func configGoAdminLanguageByYaml() string {
	langConf := viper.GetString("go_admin.language")
	var langRes = language.CN
	switch langConf {
	case "CN":
		langRes = language.CN
	case "TC":
		langRes = language.TC
	case "EN":
		langRes = language.EN
	case "JP":
		langRes = language.JP
	default:
		// not set
	}
	return langRes
}
