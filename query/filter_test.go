package query

import "testing"

func TestReadFilters(t *testing.T) {
	type TestCase struct {
		Input  string
		Opt    *ReadFilterOptions
		Output []Filter
		Err    error
	}

	testCases := []TestCase{
		{Input: ""},
		{
			Input: "filter=title eq Spaghetti",
			Output: []Filter{
				{Field: "title", Operator: "eq", Value: "Spaghetti"},
			},
		},
		{
			Input: "filter=title eq Bolognese&filter=serves gte 4",
			Output: []Filter{
				{Field: "title", Operator: "eq", Value: "Bolognese"},
				{Field: "serves", Operator: "gte", Value: "4"},
			},
		},
	}

	for _, tc := range testCases {
		t.Logf("Testing %q with options %+v", tc.Input, tc.Opt)

		filters, err := ReadStringFilters(tc.Input, nil)

		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
			break
		}

		if tc.Err != nil {
			break
		}

		if tc.Output == nil && filters != nil {
			t.Error("Expected nil")
			break
		}
		if len(filters) != len(tc.Output) {
			t.Errorf("Expected %d filters, got %d", len(tc.Output), len(filters))
		}

		for i, filter := range tc.Output {
			if i == len(filters) {
				break
			}
			if filter != filters[i] {
				t.Errorf("Expected %+v for filter %d, got %+v", filter, i, filters[i])
			}
		}
	}
}
