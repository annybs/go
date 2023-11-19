package validate

import (
	"strings"
)

// Validation error.
var (
	ErrDisallowedChars = NewError("contains disallowed characters")
)

// Chars validates whether a string contains only allowed characters.
func Chars(allow string) func(string) error {
	return func(value string) error {
		for _, r := range value {
			if !strings.ContainsRune(allow, r) {
				return ErrDisallowedChars
			}
		}
		return nil
	}
}
