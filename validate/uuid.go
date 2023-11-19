package validate

import (
	"regexp"
)

// Validation error.
var (
	ErrInvalidUUID = NewError("invalid UUID")
)

var uuidRegexp = regexp.MustCompile("^[a-f0-9]{8}(-[a-f0-9]{4}){3}-[a-f0-9]{12}$")

// UUID validates a UUID string.
// The UUID must be formatted with separators.
func UUID(value string) error {
	if !uuidRegexp.MatchString(value) {
		return ErrInvalidUUID
	}

	return nil
}
