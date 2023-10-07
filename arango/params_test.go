package arango

import "testing"

func TestReadParams(t *testing.T) {
	type TestCase struct {
		Input  string
		Output []string
	}

	testCases := []TestCase{
		{
			Input:  "FOR doc IN @@collection FILTER doc.title == @title RETURN doc",
			Output: []string{"collection", "title"},
		},
	}

	for _, tc := range testCases {
		t.Log("Input:", tc.Input)
		t.Log("Expected output:", tc.Output)

		params := ReadParams(tc.Input)

		if len(params) != len(tc.Output) {
			t.Errorf("Expected %d parameters", len(tc.Output))
			break
		}

		for i, name := range tc.Output {
			if params[i] != name {
				t.Errorf("Expected parameter %d to be %q", i, name)
			}
		}
	}
}
