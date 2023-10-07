package validate

import (
	"errors"
)

// Validation error.
var (
	ErrTooFewChars  = errors.New("too few characters")
	ErrTooManyChars = errors.New("too many characters")
)

// MaxLength validates the length of a string as being less than or equal to a given maximum.
func MaxLength(l int) func(string) error {
	return func(value string) error {
		if len(value) > l {
			return ErrTooManyChars
		}
		return nil
	}
}

// MinLength validates the length of a string as being greater than or equal to a given minimum.
func MinLength(l int) func(string) error {
	return func(value string) error {
		if len(value) < l {
			return ErrTooFewChars
		}
		return nil
	}
}
