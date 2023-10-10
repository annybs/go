package random

import "math/rand"

// Str generates a random string of the specified length using an alphanumeric character set.
var Str = String([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))

// String returns a function that can be used to generate a random string of the specified length using the given character set.
func String(chars []rune) func(length int) string {
	max := len(chars)
	return func(length int) string {
		s := make([]rune, length)
		for i := range s {
			s[i] = chars[rand.Intn(max)]
		}
		return string(s)
	}
}
