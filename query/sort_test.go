package query

import "testing"

func TestReadSorts(t *testing.T) {
	type TestCase struct {
		Input  string
		Opt    *ReadSortOptions
		Output []Sort
		Err    error
	}

	testCases := []TestCase{
		{Input: ""},
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
		t.Logf("Testing %q with options %+v", tc.Input, tc.Opt)

		sorts, err := ReadStringSorts(tc.Input, nil)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
			break
		}

		if tc.Err != nil {
			break
		}

		if tc.Output == nil && sorts != nil {
			t.Error("Expected nil")
			break
		}
		if len(sorts) != len(tc.Output) {
			t.Errorf("Expected %d sorts, got %d", len(tc.Output), len(sorts))
		}

		for i, sort := range tc.Output {
			if i == len(sorts) {
				break
			}
			if sort != sorts[i] {
				t.Errorf("Expected %+v for sort %d, got %+v", sort, i, sorts[i])
			}
		}
	}
}
