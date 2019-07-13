package database

func (receiver Type) Validate() error {
	if Nothing() == receiver {
		return internalNotLoaded{"database: not loaded"}
	}
	if receiver.errored() {
		return receiver.err
	}

	return nil
}
