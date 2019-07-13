package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type Type struct {
	db    *sql.DB
	err    error
	loaded bool
}

func (receiver Type) errored() bool {
	return nil != receiver.err
}

func Nothing() Type {
	return Type{}
}

func Something(db *sql.DB) Type {
	return Type{
		loaded: true,
		db: db,
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
