package arango

import (
	"errors"
)

// Operator error.
var (
	ErrInvalidSortDirection = errors.New("invalid sort direction")
)

var (
	sorts = map[string]string{
		// Idempotent
		"ASC":  "ASC",
		"DESC": "DESC",

		// Compatible with Sort.Direction in github.com/annybs/go/qs
		//
		// Although lowercase keywords can be used in AQL, uppercase is favoured for stylistic consistency.
		"asc":  "ASC",
		"desc": "DESC",
	}
)

// ParseSortDirection returns the valid AQL operator for an arbitrary direction string.
// This supports different inputs, such as Sort.Direction in github.com/annybs/go/qs
//
// If the input operator cannot be mapped to AQL, this function returns ErrInvalidSortDirection.
func ParseSortDirection(op string) (string, error) {
	if sorts[op] == "" {
		return "", ErrInvalidSortDirection
	}
	return sorts[op], nil
}
