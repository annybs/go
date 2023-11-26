package validate

import "net/url"

// Validation error.
var (
	ErrInvalidURL Error = NewError("invalid URL")
)

// URL validates a URL.
func URL(value string) error {
	if _, err := url.ParseRequestURI(value); err != nil {
		return ErrInvalidURL
	}

	return nil
}
