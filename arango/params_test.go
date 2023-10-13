package arango

import "testing"

func TestReadParams(t *testing.T) {
	type TestCase struct {
		Input  string
		Output []string
	}

	testCases := []TestCase{
		{
			Input: "FOR doc IN recipes FILTER doc.title == \"Spaghetti\" RETURN doc",
		},
		{
			Input:  "FOR doc IN recipes FILTER doc.title == @title RETURN doc",
			Output: []string{"title"},
		},
		{
			Input:  "FOR doc IN @@collection FILTER doc.title == @title RETURN doc",
			Output: []string{"@collection", "title"},
		},
	}

	for n, tc := range testCases {
		t.Logf("(%d) Testing %q", n, tc.Input)

		params := ReadParams(tc.Input)

		if params == nil {
			t.Errorf("Expected empty slice, got nil")
			continue
		}

		if len(params) != len(tc.Output) {
			t.Errorf("Expected %d parameters", len(tc.Output))
		}

		for i, name := range tc.Output {
			if i == len(params) {
				break
			}
			if name != params[i] {
				t.Errorf("Expected %s for parameter %d, got %s", name, i, params[i])
			}
		}
	}
}
