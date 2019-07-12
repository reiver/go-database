package database_conclusion

import (
	"fmt"
)

func (receiver Type) GoString() string {
	switch receiver {
	case Nothing():
		return "database_conclusion.Nothing()"
	case Something():
		return "database_conclusion.Something()"
	default:
		return fmt.Sprintf("database_conclusion.Error(%q)", receiver.err)
	}
}
