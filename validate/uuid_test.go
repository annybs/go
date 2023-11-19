package validate

import (
	"errors"
	"testing"
)

func TestUUID(t *testing.T) {
	type TestCase struct {
		Input string
		Err   error
	}

	testCases := []TestCase{
		{Input: "00000000-0000-0000-0000-000000000000"},
		{Input: "01234567-89ab-cdef-0123-456789abcdef"},
		{Input: "abcdef01-2345-6789-abcd-ef0123456789"},
		{Input: "Not a UUID", Err: ErrInvalidUUID},
		{Input: "00000000-00-0000-0000-00000000000000", Err: ErrInvalidUUID},
		{Input: "00000000000000000000000000000000", Err: ErrInvalidUUID},
		{Input: "01234567-89ab-cdef-ghij-klmnopqrstuv", Err: ErrInvalidUUID},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q", n, tc.Input)

		err := UUID(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
