package hash

import "testing"

func TestSHA256(t *testing.T) {
	type TestCase struct {
		Input  string
		Output string
	}

	testCases := []TestCase{
		{Input: "Gnocchi", Output: "5590208af75dc079e5e500fa39c982449ecbb02010f093a4de8b3377250c6f99"},
		{Input: "Spaghetti", Output: "dd81cff28075d9126e35ce2ba895a3ce95a9ecdc1ac899418ace10e7a697127d"},
	}

	for i, tc := range testCases {
		t.Logf("(%d) Testing %q", i, tc.Input)

		out := SHA256(tc.Input)
		if out != tc.Output {
			t.Errorf("Expected %q, got %q", tc.Output, out)
		}
	}
}

func TestSHA512(t *testing.T) {
	type TestCase struct {
		Input  string
		Output string
	}

	testCases := []TestCase{
		{Input: "Gnocchi", Output: "43d0c2d7ee218cc30221c83e57c4ea3e6a68ce73225aff9b77c148e86f97c0b5d98a018a503115c4a81b9364cf132f5fc42878ece0412346aa60f2ea01f54756"},
		{Input: "Spaghetti", Output: "259f2043005a92915be7bd288050cffdc4f55201c823b5837e22b053ef982e364bf3625ee41c676b91047a450b33e66660ec6346690b234c2a0c47c832043ef2"},
	}

	for i, tc := range testCases {
		t.Logf("(%d) Testing %q", i, tc.Input)

		out := SHA512(tc.Input)
		if out != tc.Output {
			t.Errorf("Expected %q, got %q", tc.Output, out)
		}
	}
}
