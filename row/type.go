package database_row

import (
	"database/sql"
	"errors"
	"fmt"
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

func Something(row *sql.Row) Type {
	return Type{
		loaded: true,
		row: row,
	}
}

func Err(err error) Type {
	return Type{
		loaded: true,
		err: err,
	}
}

func Error(s string) Type {
	err := errors.New(s)
	return Err(err)
}

func Errorf(format string, a ...interface{}) Type {
	err := fmt.Errorf(format, a...)
	return Err(err)
}

func ErrorNotLoaded(s string) Type {
	err := internalNotLoaded{s}
	return Err(err)
}
