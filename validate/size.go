package validate

import (
	"errors"
)

// Validation error.
var (
	ErrTooFewItems  = errors.New("too few items")
	ErrTooManyItems = errors.New("too many items")
)

// MaxSize validates the length of a slice as being less than or equal to a given maximum.
func MaxSize[T any](l int) func([]T) error {
	return func(value []T) error {
		if len(value) > l {
			return ErrTooManyItems
		}
		return nil
	}
}

// MinSize validates the length of a slice as being greater than or equal to a given minimum.
func MinSize[T any](l int) func([]T) error {
	return func(value []T) error {
		if len(value) < l {
			return ErrTooFewItems
		}
		return nil
	}
}
