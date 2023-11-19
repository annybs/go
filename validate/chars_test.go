package validate

import (
	"errors"
	"testing"
)

func TestChars(t *testing.T) {
	type TestCase struct {
		Input string
		C     string
		Err   error
	}

	hexRange := "0123456789abcdef"

	testCases := []TestCase{
		{Input: "abcd1234", C: hexRange},
		{Input: "abcd 1234", C: hexRange, Err: ErrDisallowedChars},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q against %q", n, tc.Input, tc.C)

		f := Chars(tc.C)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
