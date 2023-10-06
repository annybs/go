package validate

import "testing"

func TestInvalidUUID(t *testing.T) {
	invalid := []string{
		"Not a UUID",
		"00000000-00-0000-0000-00000000000000",
		"00000000000000000000000000000000",
		"01234567-89ab-cdef-ghij-klmnopqrstuv",
	}

	for _, uuid := range invalid {
		if err := UUID(uuid); err == nil {
			t.Errorf("%s is not a valid UUID", uuid)
		}
	}
}

func TestValidUUID(t *testing.T) {
	valid := []string{
		"00000000-0000-0000-0000-000000000000",
		"01234567-89ab-cdef-0123-456789abcdef",
		"abcdef01-2345-6789-abcd-ef0123456789",
	}

	for _, uuid := range valid {
		if err := UUID(uuid); err != nil {
			t.Errorf("%s is a valid UUID", uuid)
		}
	}
}
