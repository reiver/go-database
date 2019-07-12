package database_conclusion

type NotLoaded interface {
	error
	NotLoaded()
}

type internalNotLoaded struct {
	msg string
}

func (receiver internalNotLoaded) Error() string {
	return receiver.msg
}

func (internalNotLoaded) NotLoaded() {
	// Nothing here.
}
