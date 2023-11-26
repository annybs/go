package validate

import (
	"errors"
	"testing"
)

func TestURL(t *testing.T) {
	type TestCase struct {
		Input string
		Err   error
	}

	testCases := []TestCase{
		{Input: "http://example.com"},
		{Input: "http://subdomain.example.com"},
		{Input: "http://www.example.com/some-page.html"},
		{Input: "subdomain.com", Err: ErrInvalidURL},
		{Input: "not a url", Err: ErrInvalidURL},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q", n, tc.Input)

		err := URL(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
