package database_row

import (
	"github.com/reiver/go-database/conclusion"
)

func (receiver Type) Scan(dest ...interface{}) database_conclusion.Type {
	if Nothing() == receiver {
		return database_conclusion.ErrorNotLoaded("database_row: not loaded")
	}
	if receiver.errored() {
		return database_conclusion.Err(receiver.err)
	}

	row := receiver.row
	if nil == row {
		return database_conclusion.Error("database_row: internal error: nil row")
	}

	if err := row.Scan(dest...); nil != err {
		return database_conclusion.Err(err)
	}

	return database_conclusion.Something()
}
