package database

import (
	"github.com/reiver/go-database/conclusion"
)

func (receiver Type) Ping() database_conclusion.Type {
	if Nothing() == receiver {
		return database_conclusion.ErrorNotLoaded("database: not loaded")
	}
	if receiver.errored() {
		return database_conclusion.Err(receiver.err)
	}

	db := receiver.db
	if nil == db {
		return database_conclusion.Error("database: internal error: nil db")
	}

	if err := db.Ping(); nil != err {
		return database_conclusion.Err(err)
	}

	return database_conclusion.Something()
}
