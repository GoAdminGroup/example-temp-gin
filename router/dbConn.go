package router

import (
	"fmt"
	"github.com/GoAdminGroup/example-temp-gin/model/dbglobal"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/spf13/viper"
)

func initDBConnection(eng *engine.Engine) error {
	dbType := viper.GetString("goAdmin.dataBases.default.driver")
	var dbConn db.Connection
	switch dbType {
	default:
		return fmt.Errorf("unknown type of db driver: %v", dbType)
	case "mysql":
		dbConn = eng.MysqlConnection()
	case "sqlite":
		dbConn = eng.SqliteConnection()
	case "postgresql":
		dbConn = eng.PostgresqlConnection()
	case "mssql":
		dbConn = eng.MssqlConnection()
	}
	dbglobal.SetGlobalConn(dbConn)
	return nil
}
