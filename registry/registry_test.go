package registry

import (
	"errors"
	"testing"
)

func TestRegistry(t *testing.T) {
	type TestCase struct {
		Key   string
		Value string
		Lock  bool
		Err   error
	}

	testCases := []TestCase{
		{Key: "abc", Value: "123"},
		{Key: "abc", Value: "456"},
		{Key: "def", Value: "789", Lock: true},
		{Key: "def", Value: "000", Err: ErrKeyLocked},
	}

	r := New()
	for i, tc := range testCases {
		t.Logf("(%d) Testing set %s=%q (lock: %t)", i, tc.Key, tc.Value, tc.Lock)

		err := r.Set(tc.Key, tc.Value)
		if !errors.Is(err, tc.Err) {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}

		value := r.Get(tc.Key)
		if value != tc.Value {
			if err == nil {
				t.Errorf("Expected value %q, got %q", tc.Value, value)
			}
		}

		if tc.Lock {
			r.Lock(tc.Key)
		}
	}
}
