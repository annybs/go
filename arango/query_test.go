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
		// Append with parameters
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
		// Append with too many parameters
		{
			Input: NewQuery().
				Append("FOR doc IN @@collection", "recipes", "ignored").
				Append("FILTER doc.title == @title", "Spaghetti", "also ignored").
				Append("RETURN doc"),
			ExpectedStr: `FOR doc IN @@collection
FILTER doc.title == @title
RETURN doc`,
			ExpectedParams: map[string]any{
				"collection": "recipes",
				"title":      "Spaghetti",
			},
		},
		// Append and bind
		{
			Input: NewQuery().
				Append("FOR doc IN @@collection").
				Append("FILTER doc.title == @title").
				Append("RETURN doc").
				Bind("collection", "recipes").
				Bind("title", "Spaghetti"),
			ExpectedStr: `FOR doc IN @@collection
FILTER doc.title == @title
RETURN doc`,
			ExpectedParams: map[string]any{
				"collection": "recipes",
				"title":      "Spaghetti",
			},
		},
		// Append and bind map
		{
			Input: NewQuery().
				Append("FOR doc IN @@collection").
				Append("FILTER doc.title == @title").
				Append("RETURN doc").
				BindMap(map[string]any{"collection": "recipes", "title": "Spaghetti"}),
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
		t.Logf("Testing %+v", tc.Input)

		actualStr := tc.Input.String()

		if actualStr != tc.ExpectedStr {
			t.Errorf("Expected %q, got %q", tc.ExpectedStr, actualStr)
		}

		if len(tc.Input.Params) != len(tc.ExpectedParams) {
			t.Errorf("Expected %d parameters, got %d", len(tc.ExpectedParams), len(tc.Input.Params))
		}

		for name, value := range tc.ExpectedParams {
			if tc.Input.Params[name] != value {
				t.Errorf("Expected parameter %q to be %v; got %v", name, value, tc.Input.Params[name])
			}
		}
	}
}
