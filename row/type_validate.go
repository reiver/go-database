package database_row

func (receiver Type) Validate() error {
	if Nothing() == receiver {
		return internalNotLoaded{"database_row: not loaded"}
	}
	if receiver.errored() {
		return receiver.err
	}

	return nil
}
