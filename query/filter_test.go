package query

import "testing"

func TestReadFilters(t *testing.T) {
	type TestCase struct {
		Input  string
		Output []Filter
		Err    bool
	}

	testCases := []TestCase{
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
		t.Logf("%q", tc.Input)

		filters, err := ReadStringFilters(tc.Input, nil)
		if tc.Err {
			if err == nil {
				t.Fatal("Expected error; got nil")
			}
		} else {
			if err != nil {
				t.Fatalf("Expected no error; got %s", err)
			}
		}

		if len(filters) != len(tc.Output) {
			t.Errorf("Expected %d filters, got %d", len(tc.Output), len(filters))
		}
		for i, filter := range tc.Output {
			if i == len(filters) {
				break
			}
			if filter != filters[i] {
				t.Errorf("Expected filter %d to be %q, got %q", i, filter, filters[i])
			}
		}
	}
}
