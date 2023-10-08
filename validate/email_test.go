package validate

import "testing"

func TestEmail(t *testing.T) {
	type TestCase struct {
		Input string
		Err   error
	}

	testCases := []TestCase{
		{Input: "test@example.com"},
		{Input: "testexample.com", Err: ErrInvalidEmail},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q", n, tc.Input)

		err := Email(tc.Input)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
