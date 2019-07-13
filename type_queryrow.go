package database

import (
	"github.com/reiver/go-database/row"

	"context"
)

func (receiver Type) QueryRow(ctx context.Context, query string, args ...interface{}) database_row.Type {
	if Nothing() == receiver {
		return database_row.ErrorNotLoaded("database: not loaded")
	}
	if receiver.errored() {
		return database_row.Err(receiver.err)
	}

	db := receiver.db
	if nil == db {
		return database_row.Error("database_transaction: internal error: nil db")
	}

	row := db.QueryRowContext(ctx, query, args...)

	return database_row.Something(row)
}
