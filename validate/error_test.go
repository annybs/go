package validate

import (
	"errors"
	"testing"
)

func TestErrorIs(t *testing.T) {
	type TestCase struct {
		Err    error
		Target error
		Is     bool
	}

	testCases := []TestCase{
		// Is any validation error
		{Err: Err, Target: Err, Is: true},
		{Err: ErrDisallowedChars, Target: Err, Is: true},
		{Err: ErrMustBeGreater, Target: Err, Is: true},

		// Is specific validation error
		{Err: ErrDisallowedChars, Target: ErrDisallowedChars, Is: true},
		{Err: ErrMustBeGreater, Target: ErrMustBeGreater, Is: true},

		// Is not specific validation error
		{Err: Err, Target: ErrDisallowedChars},
		{Err: Err, Target: ErrMustBeGreater},
		{Err: ErrMustBeGreater, Target: ErrDisallowedChars},
		{Err: ErrDisallowedChars, Target: ErrMustBeGreater},

		// Is not any other error
		{Err: ErrDisallowedChars, Target: errors.New("contains disallowed characters")},
		{Err: ErrMustBeGreater, Target: errors.New("must be greater than %v")},
	}

	for i, tc := range testCases {
		t.Logf("(%d) Testing %v against %v", i, tc.Err, tc.Target)

		if errors.Is(tc.Err, tc.Target) {
			if !tc.Is {
				t.Errorf("%v should not equal %v", tc.Err, tc.Target)
			}
		} else {
			if tc.Is {
				t.Errorf("%v should equal %v", tc.Err, tc.Target)
			}
		}
	}
}
