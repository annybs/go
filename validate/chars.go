package validate

import (
	"fmt"
	"strings"
)

// Chars validates whether a string contains only allowed characters.
func Chars(allow string) func(string) error {
	return func(value string) error {
		rs := []rune(value)
		for _, r := range rs {
			if !strings.ContainsRune(allow, r) {
				return fmt.Errorf("Contains disallowed characters")
			}
		}
		return nil
	}
}
