package database

import (
	"database/sql"
)

func Open(driverName string, dataSourceName string) Type {

	db, err := sql.Open(driverName, dataSourceName)
	if nil != err {
		return Err(err)
	}

	return Something(db)
}
