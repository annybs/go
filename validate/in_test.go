package validate

import (
	"testing"
)

func TestInString(t *testing.T) {
	type TestCase struct {
		Input string
		A     []string
		Err   error
	}

	allow := []string{"abcd", "ef", "1234"}
	testCases := []TestCase{
		{Input: "abcd", A: allow},
		{Input: "ef", A: allow},
		{Input: "1234", A: allow},
		{Input: "5678", A: allow, Err: ErrValueNotAllowed},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q against %v", n, tc.Input, tc.A)

		f := In(tc.A...)
		err := f(tc.Input)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestInInt(t *testing.T) {
	type TestCase struct {
		Input int
		A     []int
		Err   error
	}

	allow := []int{1, 23, 456}
	testCases := []TestCase{
		{Input: 1, A: allow},
		{Input: 23, A: allow},
		{Input: 456, A: allow},
		{Input: 789, A: allow, Err: ErrValueNotAllowed},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %d against %v", n, tc.Input, tc.A)

		f := In(tc.A...)
		err := f(tc.Input)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
