package arango

import (
	"testing"
)

func TestQueryAppend(t *testing.T) {
	type TestCase struct {
		Input          *Query
		ExpectedStr    string
		ExpectedParams map[string]any
	}

	testCases := []TestCase{
		{
			Input: NewQuery().
				Append("FOR doc IN @@collection", "recipes").
				Append("FILTER doc.title == @title", "Spaghetti").
				Append("RETURN doc"),
			ExpectedStr: `FOR doc IN @@collection
FILTER doc.title == @title
RETURN doc`,
			ExpectedParams: map[string]any{
				"collection": "recipes",
				"title":      "Spaghetti",
			},
		},
	}

	for _, tc := range testCases {
		actualStr := tc.Input.String()

		if actualStr != tc.ExpectedStr {
			t.Logf("Expected: %q", tc.ExpectedStr)
			t.Logf("Actual: %q", actualStr)
			t.Fail()
		}

		if len(tc.Input.Params) != len(tc.ExpectedParams) {
			t.Errorf("Expected %d parameters; got %d", len(tc.ExpectedParams), len(tc.Input.Params))
		}

		for name, value := range tc.ExpectedParams {
			if tc.Input.Params[name] == nil {
				t.Errorf("Expected parameter %q to be %q; got nil", name, value)
			} else if tc.Input.Params[name] != value {
				t.Errorf("Expected parameter %q to be %q; got %q", name, value, tc.Input.Params[name])
			}
		}
	}
}
