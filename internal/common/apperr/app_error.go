package apperr

import (
	"errors"
	"fmt"
)

var ErrApp = errors.New("")

func New(message string) error {
	return fmt.Errorf(message+"%w", ErrApp)
}
