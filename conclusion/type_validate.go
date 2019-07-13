package database_conclusion

func (receiver Type) Validate() error {
	switch receiver {
	case Nothing():
		return internalNotLoaded{"database_conclusion: not loaded"}
	case Something():
		return nil
	default:
		return receiver.err
	}
}
