package dbglobal

import "github.com/GoAdminGroup/go-admin/modules/db"

var globalConn db.Connection

func SetGlobalConn(conn db.Connection) {
	globalConn = conn
}
func connection() *db.SQL {
	return db.WithDriver(globalConn)
}

func Table(table string) *db.SQL {
	return connection().Table(table)
}

// most of use is
//	item, _ := dbglobal.TableItem("table_name", id)
func TableItem(table string, args interface{}) (map[string]interface{}, error) {
	return Table(table).Find(args)
}
