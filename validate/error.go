package validate

import "fmt"

// Validation error.
var (
	Err = Error{}
)

// Error represents a validation error.
type Error struct {
	Message string
	Data    []any
}

// Error retrieves the message of a validation Error.
// If it has Data, the message will be formatted.
func (e Error) Error() string {
	if len(e.Data) > 0 {
		return fmt.Sprintf(e.Message, e.Data...)
	}
	return e.Message
}

// Is determines whether the Error is an instance of the target.
// https://pkg.go.dev/errors#Is
//
// If the target is a validation error and specifies a message, this function returns true if the messages match.
// If the target is an empty validation error, this function always returns true.
func (e Error) Is(target error) bool {
	if t, ok := target.(Error); ok {
		return t.Message == e.Message || t.Message == ""
	}
	return false
}

func (e Error) With(value any) Error {
	if e.Data == nil {
		e.Data = []any{}
	}
	e.Data = append(e.Data, value)
	return e
}

// NewError creates a new validation error.
func NewError(message string) Error {
	return Error{
		Message: message,
	}
}
