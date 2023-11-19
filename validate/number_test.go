package validate

import (
	"errors"
	"testing"
)

func TestMax(t *testing.T) {
	type TestCase struct {
		Input int
		N     int
		Excl  bool
		Err   error
	}

	testCases := []TestCase{
		{Input: 10, N: 0, Err: ErrMustBeLessOrEqual},
		{Input: 10, N: 10},
		{Input: 10, N: 15},
		{Input: 10, N: 10, Excl: true, Err: ErrMustBeLess},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %d against maximum of %d", n, tc.Input, tc.N)

		f := Max(tc.N, tc.Excl)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestMaxFloat32(t *testing.T) {
	type TestCase struct {
		Input float32
		N     float32
		Excl  bool
		Err   error
	}

	testCases := []TestCase{
		{Input: 10, N: 0, Err: ErrMustBeLessOrEqual},
		{Input: 10, N: 10},
		{Input: 10, N: 15},
		{Input: 10, N: 10, Excl: true, Err: ErrMustBeLess},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %g against maximum of %g", n, tc.Input, tc.N)

		f := MaxFloat32(tc.N, tc.Excl)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestMaxFloat64(t *testing.T) {
	type TestCase struct {
		Input float64
		N     float64
		Excl  bool
		Err   error
	}

	testCases := []TestCase{
		{Input: 10, N: 0, Err: ErrMustBeLessOrEqual},
		{Input: 10, N: 10},
		{Input: 10, N: 15},
		{Input: 10, N: 10, Excl: true, Err: ErrMustBeLess},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %g against maximum of %g", n, tc.Input, tc.N)

		f := MaxFloat64(tc.N, tc.Excl)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestMin(t *testing.T) {
	type TestCase struct {
		Input int
		N     int
		Excl  bool
		Err   error
	}

	testCases := []TestCase{
		{Input: 10, N: 0},
		{Input: 10, N: 10},
		{Input: 10, N: 15, Err: ErrMustBeGreaterOrEqual},
		{Input: 10, N: 10, Excl: true, Err: ErrMustBeGreater},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %d against minimum of %d", n, tc.Input, tc.N)

		f := Min(tc.N, tc.Excl)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestMinFloat32(t *testing.T) {
	type TestCase struct {
		Input float32
		N     float32
		Excl  bool
		Err   error
	}

	testCases := []TestCase{
		{Input: 10, N: 0},
		{Input: 10, N: 10},
		{Input: 10, N: 15, Err: ErrMustBeGreaterOrEqual},
		{Input: 10, N: 10, Excl: true, Err: ErrMustBeGreater},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %g against minimum of %g", n, tc.Input, tc.N)

		f := MinFloat32(tc.N, tc.Excl)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestMinFloat64(t *testing.T) {
	type TestCase struct {
		Input float64
		N     float64
		Excl  bool
		Err   error
	}

	testCases := []TestCase{
		{Input: 10, N: 0},
		{Input: 10, N: 10},
		{Input: 10, N: 15, Err: ErrMustBeGreaterOrEqual},
		{Input: 10, N: 10, Excl: true, Err: ErrMustBeGreater},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %g against minimum of %g", n, tc.Input, tc.N)

		f := MinFloat64(tc.N, tc.Excl)
		err := f(tc.Input)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
