package validate

import (
	"errors"
	"testing"
)

func TestMaxSize(t *testing.T) {
	type TestCase struct {
		Input []int
		L     int
		Err   error
	}

	testCases := []TestCase{
		{Input: []int{1, 2, 3, 4}, L: 8},
		{Input: []int{1, 2, 3, 4, 5, 6, 7, 8}, L: 8},
		{Input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, L: 8, Err: ErrMustHaveFewerItems},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q against maximum length of %d", n, tc.Input, tc.L)

		f := MaxSize[int](tc.L)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestMinSize(t *testing.T) {
	type TestCase struct {
		Input []int
		L     int
		Err   error
	}

	testCases := []TestCase{
		{Input: []int{1, 2, 3, 4}, L: 8, Err: ErrMustHaveMoreItems},
		{Input: []int{1, 2, 3, 4, 5, 6, 7, 8}, L: 8},
		{Input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, L: 8},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q against minimum length of %d", n, tc.Input, tc.L)

		f := MinSize[int](tc.L)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
