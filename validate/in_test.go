package validate

import (
	"testing"
)

func TestIn(t *testing.T) {
	type TestCase[T comparable] struct {
		Input T
		A     []T
		Err   error
	}

	strIn := []string{"abcd", "ef", "1234"}
	strTestCases := []TestCase[string]{
		{Input: "abcd", A: strIn},
		{Input: "ef", A: strIn},
		{Input: "1234", A: strIn},
		{Input: "5678", A: strIn, Err: ErrValueNotAllowed},
	}

	for _, tc := range strTestCases {
		t.Logf("Testing %q against %v", tc.Input, tc.A)

		f := In(tc.A...)
		err := f(tc.Input)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}

	intIn := []int{1, 23, 456}
	intTestCases := []TestCase[int]{
		{Input: 1, A: intIn},
		{Input: 23, A: intIn},
		{Input: 456, A: intIn},
		{Input: 789, A: intIn, Err: ErrValueNotAllowed},
	}

	for _, tc := range intTestCases {
		t.Logf("Testing %d against %v", tc.Input, tc.A)

		f := In(tc.A...)
		err := f(tc.Input)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
