package validate

import "testing"

func TestChars(t *testing.T) {
	type TestCase struct {
		C     string
		Input string
		Err   bool
	}

	hexRange := "0123456789abcdef"

	testCases := []TestCase{
		{C: hexRange, Input: "abcd1234"},
		{C: hexRange, Input: "abcd 1234", Err: true},
	}

	for _, tc := range testCases {
		t.Logf("%q contains only allowed characters %q", tc.Input, tc.C)

		f := Chars(tc.C)
		err := f(tc.Input)

		if tc.Err {
			if err == nil {
				t.Error("Expected error; got nil")
			}
		} else {
			if err != nil {
				t.Errorf("Expected nil; got %s", err)
			}
		}
	}
}
