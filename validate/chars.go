package validate

import (
	"errors"
	"strings"
)

// Validation error.
var (
	ErrDisallowedChars = errors.New("contains disallowed characters")
)

// Chars validates whether a string contains only allowed characters.
func Chars(allow string) func(string) error {
	return func(value string) error {
		rs := []rune(value)
		for _, r := range rs {
			if !strings.ContainsRune(allow, r) {
				return ErrDisallowedChars
			}
		}
		return nil
	}
}
