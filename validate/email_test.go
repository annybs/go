package validate

import "testing"

func TestEmail(t *testing.T) {
	type TestCase struct {
		Input string
		Err   bool
	}

	testCases := []TestCase{
		{Input: "test@example.com"},
		{Input: "testexample.com", Err: true},
	}

	for _, tc := range testCases {
		err := Email(tc.Input)
		if tc.Err {
			if err == nil {
				t.Errorf("Expected %q to be an invalid email; got nil", tc.Input)
			}
		} else {
			if err != nil {
				t.Errorf("Expected %q to be a valid email; got %s", tc.Input, err)
			}
		}
	}
}
