package dbglobal

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
)

var globalConn *db.Connection

func Conn() db.Connection {
	if globalConn == nil {
		return nil
	}
	return *globalConn
}

func SetGlobalConn(conn db.Connection) {
	if globalConn == nil {
		globalConn = &conn
	}
}

func connection() *db.SQL {
	if globalConn == nil {
		return nil
	}
	return db.WithDriver(*globalConn)
}

func Table(table string) (*db.SQL, error) {
	if connection() == nil {
		return nil, fmt.Errorf("connect is empty")
	}
	return connection().Table(table), nil
}

// find item by id
//	item, _ := dbglobal.TableItem("table_name", id)
func TableItem(table string, args interface{}) (map[string]interface{}, error) {
	sql, err := Table(table)
	if err != nil {
		return nil, err
	}
	return sql.Find(args)
}

func TableInsert(table string, values dialect.H) (int64, error) {
	if connection() == nil {
		return 0, fmt.Errorf("connect is empty")
	}
	sql := connection().Table(table)
	return sql.Insert(values)
}

func TableUpdateByID(table string, id int64, values dialect.H) (int64, error) {
	if connection() == nil {
		return 0, nil
	}
	sql := connection().Table(table)
	return sql.Where("id", "=", id).Update(values)
}

func TableCount(table string) (int64, error) {
	if connection() == nil {
		return 0, nil
	}
	sql := connection().Table(table)
	return sql.
		Count()
}

func TableCountByTimeRange(table string, timeField string, startTimestamp, endTimestamp string) (int64, error) {
	if connection() == nil {
		return 0, nil
	}
	sql := connection().Table(table)
	return sql.
		Select(timeField).
		Where(timeField, ">=", startTimestamp).
		Where(timeField, "<=", endTimestamp).
		Count()
}

func TableCountByKey(table string, key string, arg interface{}) (int64, error) {
	if connection() == nil {
		return 0, fmt.Errorf("connect is empty")
	}
	sql := connection().Table(table)
	return sql.Where(key, "=", arg).Count()
}

func TableFirstByKey(table string, key string, arg interface{}) (map[string]interface{}, error) {
	if connection() == nil {
		return nil, fmt.Errorf("connect is empty")
	}
	sql := connection().Table(table)
	return sql.Where(key, "=", arg).First()
}

func TableFindIDByUniqueKey(table, uniqueKey, uniqueValue string) (int64, error) {
	if connection() == nil {
		return 0, fmt.Errorf("connect is empty")
	}
	sql := connection().Table(table)
	first, err := sql.Where(uniqueKey, "=", uniqueValue).First()
	if err != nil {
		return 0, fmt.Errorf("connect is empty")
	}
	return first["id"].(int64), nil
}
