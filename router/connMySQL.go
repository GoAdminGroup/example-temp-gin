package router

import "github.com/spf13/viper"

const (
	envKeyMysqlHost   = "DB_MYSQL_HOST"
	envKeyMysqlPort   = "DB_MYSQL_PORT"
	envKeyMysqlDBName = "DB_MYSQL_DBNAME"
	envKeyMysqlDBUser = "DB_MYSQL_USER"
	envKeyMysqlDBPWD  = "DB_MYSQL_PWD"
)

func parseEnvStringOrDefault(envKey, defaultStr string) string {
	//config.ProbeViperString(
	//	"go_admin.data_bases.default.driver",
	//	"go_admin.data_bases.default.host",
	//	"go_admin.data_bases.default.port",
	//	"go_admin.data_bases.default.name",
	//	"go_admin.data_bases.default.user",
	//	"go_admin.data_bases.default.pwd",
	//	"go_admin.data_bases.default.max_idle_con",
	//	"go_admin.data_bases.default.max_open_con",
	//)
	var result string
	if viper.GetString(envKey) == "" {
		result = defaultStr
	} else {
		result = viper.GetString(envKey)
	}
	return result
}
