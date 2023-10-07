package validate

import "testing"

func TestUUID(t *testing.T) {
	type TestCase struct {
		Input string
		Err   bool
	}

	testCases := []TestCase{
		{Input: "00000000-0000-0000-0000-000000000000"},
		{Input: "01234567-89ab-cdef-0123-456789abcdef"},
		{Input: "abcdef01-2345-6789-abcd-ef0123456789"},
		{Input: "Not a UUID", Err: true},
		{Input: "00000000-00-0000-0000-00000000000000", Err: true},
		{Input: "00000000000000000000000000000000", Err: true},
		{Input: "01234567-89ab-cdef-ghij-klmnopqrstuv", Err: true},
	}

	for _, testCase := range testCases {
		err := UUID(testCase.Input)
		if testCase.Err {
			if err == nil {
				t.Errorf("Expected %q to be an invalid UUID; got nil", testCase.Input)
			}
		} else {
			if err != nil {
				t.Errorf("Expected %q to be a valid UUID; got %s", testCase.Input, err)
			}
		}
	}
}
