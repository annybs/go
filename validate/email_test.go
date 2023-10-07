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

	for _, testCase := range testCases {
		err := Email(testCase.Input)
		if testCase.Err {
			if err == nil {
				t.Errorf("Expected %q to be an invalid email; got nil", testCase.Input)
			}
		} else {
			if err != nil {
				t.Errorf("Expected %q to be a valid email; got %s", testCase.Input, err)
			}
		}
	}
}
