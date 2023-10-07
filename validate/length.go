package validate

import (
	"errors"
	"fmt"
)

// MaxLength validates the length of a string as being less than or equal to a given maximum.
func MaxLength(l int) func(string) error {
	return func(value string) error {
		if len(value) > l {
			if l != 1 {
				return fmt.Errorf("Must not be longer than %d characters", l)
			}
			return errors.New("Must not be longer than 1 character")
		}
		return nil
	}
}

// MinLength validates the length of a string as being greater than or equal to a given minimum.
func MinLength(l int) func(string) error {
	return func(value string) error {
		if len(value) < l {
			if l != 1 {
				return fmt.Errorf("Must not be shorter than %d characters", l)
			}
			return errors.New("Must not be shorter than 1 character")
		}
		return nil
	}
}
