package query

import "testing"

func TestReadSorts(t *testing.T) {
	type TestCase struct {
		Input  string
		Output []Sort
		Err    bool
	}

	testCases := []TestCase{
		{
			Input: "sort=title asc",
			Output: []Sort{
				{Field: "title", Direction: "asc"},
			},
		},
		{
			Input: "sort=title asc&sort=serves asc",
			Output: []Sort{
				{Field: "title", Direction: "asc"},
				{Field: "serves", Direction: "asc"},
			},
		},
	}

	for _, tc := range testCases {
		t.Logf("%q", tc.Input)

		sorts, err := ReadStringSorts(tc.Input, nil)
		if tc.Err {
			if err == nil {
				t.Fatal("Expected error; got nil")
			}
		} else {
			if err != nil {
				t.Fatalf("Expected no error; got %s", err)
			}
		}

		if len(sorts) != len(tc.Output) {
			t.Errorf("Expected %d sorts, got %d", len(tc.Output), len(sorts))
		}
		for i, sort := range tc.Output {
			if i == len(sorts) {
				break
			}
			if sort != sorts[i] {
				t.Errorf("Expected sort %d to be %q, got %q", i, sort, sorts[i])
			}
		}
	}
}
