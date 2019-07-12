package database_conclusion

import (
	"errors"
	"fmt"
)

type Type struct {
	err    error
	loaded bool
}

func Nothing() Type {
	return Type{}
}

func Something() Type {
	return Type{
		loaded: true,
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
