package dbglobal

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
)

var globalConn db.Connection

func SetGlobalConn(conn db.Connection) {
	globalConn = conn
}
func connection() *db.SQL {
	if globalConn == nil {
		return nil
	}
	return db.WithDriver(globalConn)
}

func Table(table string) (*db.SQL, error) {
	if connection() == nil {
		return nil, fmt.Errorf("connect is empty")
	}
	return connection().Table(table), nil
}

// most of use is
//	item, _ := dbglobal.TableItem("table_name", id)
func TableItem(table string, args interface{}) (map[string]interface{}, error) {
	sql, err := Table(table)
	if err != nil {
		return nil, err
	}
	return sql.Find(args)
}
