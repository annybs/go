package validate

import (
	"errors"
	"testing"
)

func TestEqualInt(t *testing.T) {
	type TestCase struct {
		I   int
		C   int
		Err error
	}

	testCases := []TestCase{
		{I: 1, C: 1},
		{I: 5 ^ 3, C: 5 ^ 3},
		{I: 10, C: 15, Err: ErrNotEqual},
	}

	for i, tc := range testCases {
		t.Logf("(%d) Testing %d against %d", i, tc.I, tc.C)

		f := Equal(tc.C)
		err := f(tc.I)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}

func TestEqualStr(t *testing.T) {
	type TestCase struct {
		I   string
		C   string
		Err error
	}

	testCases := []TestCase{
		{I: "abc", C: "abc"},
		{I: "def ghi 123", C: "def ghi 123"},
		{I: "jkl", C: "mno", Err: ErrNotEqual},
	}

	for i, tc := range testCases {
		t.Logf("(%d) Testing %s against %s", i, tc.I, tc.C)

		f := Equal(tc.C)
		err := f(tc.I)

		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
	}
}
