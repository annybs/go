package validate

import (
	"errors"
	"testing"
)

func TestMaxLength(t *testing.T) {
	type TestCase struct {
		Input string
		L     int
		Err   error
	}

	testCases := []TestCase{
		{Input: "abcd", L: 8},
		{Input: "abcdefgh", L: 8},
		{Input: "abcd efg", L: 8},
		{Input: "abcdefghi", L: 8, Err: ErrTooManyChars},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q against maximum length of %d", n, tc.Input, tc.L)

		f := MaxLength(tc.L)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestMinLength(t *testing.T) {
	type TestCase struct {
		Input string
		L     int
		Err   error
	}

	testCases := []TestCase{
		{Input: "abcd", L: 8, Err: ErrTooFewChars},
		{Input: "abcdefgh", L: 8},
		{Input: "abcd efg", L: 8},
		{Input: "abcdefghi", L: 8},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q against minimum length of %d", n, tc.Input, tc.L)

		f := MinLength(tc.L)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
