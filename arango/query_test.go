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

	for _, testCase := range testCases {
		actualStr := testCase.Input.String()

		if actualStr != testCase.ExpectedStr {
			t.Logf("Expected: %q", testCase.ExpectedStr)
			t.Logf("Actual: %q", actualStr)
			t.Fail()
		}

		if len(testCase.Input.Params) != len(testCase.ExpectedParams) {
			t.Errorf("Expected %d parameters; got %d", len(testCase.ExpectedParams), len(testCase.Input.Params))
		}

		for name, value := range testCase.ExpectedParams {
			if testCase.Input.Params[name] == nil {
				t.Errorf("Expected parameter %q to be %q; got nil", name, value)
			} else if testCase.Input.Params[name] != value {
				t.Errorf("Expected parameter %q to be %q; got %q", name, value, testCase.Input.Params[name])
			}
		}
	}
}
