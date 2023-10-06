package validate

import (
	"errors"
	"regexp"
)

var uuidRegexp = regexp.MustCompile("^[a-f0-9]{8}(-[a-f0-9]{4}){3}-[a-f0-9]{12}$")

// UUID validates a UUID string.
// The UUID must be formatted with separators.
func UUID(value string) error {
	if !uuidRegexp.MatchString(value) {
		return errors.New("Invalid UUID")
	}

	return nil
}
