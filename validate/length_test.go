package validate

import "testing"

func TestMaxLength(t *testing.T) {
	type TestCase struct {
		L     int
		Input string
		Err   bool
	}

	testCases := []TestCase{
		{L: 8, Input: "abcd"},
		{L: 8, Input: "abcdefgh"},
		{L: 8, Input: "abcd efg"},
		{L: 8, Input: "abcdefghi", Err: true},
	}

	for _, tc := range testCases {
		t.Logf("Max length %d for %q", tc.L, tc.Input)

		f := MaxLength(tc.L)
		err := f(tc.Input)
		if tc.Err {
			if err == nil {
				t.Error("Expected an error; got nil")
			}
		} else {
			if err != nil {
				t.Errorf("Expected nil; got %s", err)
			}
		}
	}
}

func TestMinLength(t *testing.T) {
	type TestCase struct {
		L     int
		Input string
		Err   bool
	}

	testCases := []TestCase{
		{L: 8, Input: "abcd", Err: true},
		{L: 8, Input: "abcdefgh"},
		{L: 8, Input: "abcd efg"},
		{L: 8, Input: "abcdefghi"},
	}

	for _, tc := range testCases {
		t.Logf("Min length %d for %q", tc.L, tc.Input)

		f := MinLength(tc.L)
		err := f(tc.Input)
		if tc.Err {
			if err == nil {
				t.Error("Expected an error; got nil")
			}
		} else {
			if err != nil {
				t.Errorf("Expected nil; got %s", err)
			}
		}
	}
}
