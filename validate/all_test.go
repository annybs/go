package validate

import "testing"

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
		{Input: "abc", F: f, Err: ErrTooFewChars},
		{Input: "abcdef012", F: f, Err: ErrTooManyChars},
		{Input: "abcdefgh", F: f, Err: ErrDisallowedChars},
		{Input: "01abcd", F: f, Err: ErrValueNotAllowed},
	}

	for _, tc := range testCases {
		t.Logf("Testing %q", tc.Input)

		err := tc.F(tc.Input)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
