package database

import (
	"testing"
)

func TestInternalNotLoadedAsError(t *testing.T) {
	var err error = internalNotLoaded{} // THIS IS THE LINE THAT ACTUALLY MATTERS!

	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalNotLoadedANotLoaded(t *testing.T) {
	var complainer NotLoaded = internalNotLoaded{} // THIS IS THE LINE THAT ACTUALLY MATTERS!

	if nil == complainer {
		t.Error("This should never happen.")
	}
}
