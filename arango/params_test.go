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

	for _, testCase := range testCases {
		t.Log("Input:", testCase.Input)
		t.Log("Expected output:", testCase.Output)

		params := ReadParams(testCase.Input)

		if len(params) != len(testCase.Output) {
			t.Errorf("Expected %d parameters", len(testCase.Output))
			break
		}

		for i, name := range testCase.Output {
			if params[i] != name {
				t.Errorf("Expected parameter %d to be %q", i, name)
			}
		}
	}
}
