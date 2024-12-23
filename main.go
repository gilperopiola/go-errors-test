package main

import (
	"errors"
	"fmt"
)

var (
	errNormal = errors.New("normal error")
	errFunc   = func(id string) error {
		return fmt.Errorf("error on ID %s:%w", id, errNormal)
	}
	ErrNonUniqueUUID = func(UUID string) error { return fmt.Errorf("non unique UUID (%s) found in item ", UUID) }
)

func main() {

	isErrorNormal := errors.Is(errNormal, errNormal)
	isErrorFunc := errors.Is(errFunc("32"), errNormal)

	if isErrorNormal {
		println("error is normal")
	}
	if isErrorFunc {
		println("error is func")
	}

}
