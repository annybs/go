package validate

import "testing"

func TestInvalidEmail(t *testing.T) {
	valid := []string{
		"testexample.com",
	}

	for _, email := range valid {
		if err := Email(email); err == nil {
			t.Errorf("%s is not a valid email", email)
		}
	}
}

func TestValidEmail(t *testing.T) {
	valid := []string{
		"test@example.com",
	}

	for _, email := range valid {
		if err := Email(email); err != nil {
			t.Errorf("%s is a valid email", email)
		}
	}
}
