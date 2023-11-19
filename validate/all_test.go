package validate

import (
	"errors"
	"testing"
)

func TestAll(t *testing.T) {
	type TestCase[T any] struct {
		Input T
		F     func(T) error
		Err   error
	}

	f := All(
		MinLength(4),
		MaxLength(8),
		Chars("0123456789abcdef"),
		In("abcd", "abcdef", "12345678"),
	)

	testCases := []TestCase[string]{
		{Input: "abcd", F: f},
		{Input: "abcdef", F: f},
		{Input: "12345678", F: f},
		{Input: "abc", F: f, Err: ErrMustBeLonger},
		{Input: "abcdef012", F: f, Err: ErrMustBeShorter},
		{Input: "abcdefgh", F: f, Err: ErrDisallowedChars},
		{Input: "01abcd", F: f, Err: ErrValueNotAllowed},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q", n, tc.Input)

		err := tc.F(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
