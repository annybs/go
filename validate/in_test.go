package validate

import (
	"fmt"
	"strings"
	"testing"
)

func TestIn(t *testing.T) {
	type TestCase[T comparable] struct {
		S     []T
		Input T
		Err   bool
	}

	strIn := []string{"abcd", "ef", "1234"}
	strTestCases := []TestCase[string]{
		{S: strIn, Input: "abcd"},
		{S: strIn, Input: "ef"},
		{S: strIn, Input: "1234"},
		{S: strIn, Input: "5678", Err: true},
	}

	for _, tc := range strTestCases {
		t.Logf("%q in %s", tc.Input, strings.Join(tc.S, ", "))

		f := In(tc.S)
		err := f(tc.Input)

		if tc.Err {
			if err == nil {
				t.Error("Expected error; got nil")
			}
		} else {
			if err != nil {
				t.Errorf("Expected nil; got %s", err)
			}
		}
	}

	intIn := []int{1, 23, 456}
	intTestCases := []TestCase[int]{
		{S: intIn, Input: 1},
		{S: intIn, Input: 23},
		{S: intIn, Input: 456},
		{S: intIn, Input: 789, Err: true},
	}

	for _, tc := range intTestCases {
		intf := []string{}
		for _, v := range tc.S {
			intf = append(intf, fmt.Sprint(v))
		}
		t.Logf("%d in %s", tc.Input, strings.Join(intf, ", "))

		f := In(tc.S)
		err := f(tc.Input)

		if tc.Err {
			if err == nil {
				t.Error("Expected error; got nil")
			}
		} else {
			if err != nil {
				t.Errorf("Expected nil; got %s", err)
			}
		}
	}
}
