package database_row

import (
	"database/sql"
)

type Type struct {
	row    *sql.Row
	err    error
	loaded bool
}

func (receiver Type) errored() bool {
	return nil != receiver.err
}

func Nothing() Type {
	return Type{}
}

func Error(err error) Type {
	return Type{
		loaded: true,
		err: err,
	}
}

func Something(row *sql.Row) Type {
	return Type{
		loaded: true,
		row: row,
	}
}
