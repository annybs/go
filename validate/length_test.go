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

	for _, testCase := range testCases {
		t.Logf("Max length %d for %q", testCase.L, testCase.Input)

		f := MaxLength(testCase.L)
		err := f(testCase.Input)
		if testCase.Err {
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

	for _, testCase := range testCases {
		t.Logf("Min length %d for %q", testCase.L, testCase.Input)

		f := MinLength(testCase.L)
		err := f(testCase.Input)
		if testCase.Err {
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
