package arango

import (
	"errors"
)

// Operator error.
var (
	ErrInvalidOperator = errors.New("invalid operator")
)

var (
	operators = map[string]string{
		// Idempotent
		"==":       "==",
		"!=":       "!=",
		">":        ">",
		">=":       ">=",
		"<":        "<",
		"<=":       "<=",
		"IN":       "IN",
		"NOT IN":   "NOT IN",
		"LIKE":     "LIKE",
		"NOT LIKE": "NOT LIKE",

		// Compatible with Filter.Operator in github.com/recipeer/go/qs
		"eq":       "==",
		"neq":      "!=",
		"gt":       ">",
		"gte":      ">=",
		"lt":       "<",
		"lte":      "<=",
		"in":       "IN",
		"not in":   "NOT IN",
		"like":     "LIKE",
		"not like": "NOT LIKE",
	}

	arrayOperators  = []string{"IN", "NOT IN"}
	boolOperators   = []string{"==", "!="}
	numberOperators = []string{"==", "!=", ">", ">=", "<", "<="}
	stringOperators = []string{"==", "!=", ">", ">=", "<", "<=", "LIKE", "NOT LIKE"}
)

// IsArrayOperator returns true if the given operator can be used with an array value.
func IsArrayOperator(op string) bool {
	op, _ = ParseOperator(op)
	if op == "" {
		return false
	}
	for _, arrOp := range arrayOperators {
		if arrOp == op {
			return true
		}
	}
	return false
}

// IsBoolOperator returns true if the given operator can be used with a Boolean value.
func IsBoolOperator(op string) bool {
	op, _ = ParseOperator(op)
	if op == "" {
		return false
	}
	for _, boolOp := range boolOperators {
		if boolOp == op {
			return true
		}
	}
	return false
}

// IsNumberOperator returns true if the given operator can be used with a numeric value.
func IsNumberOperator(op string) bool {
	op, _ = ParseOperator(op)
	if op == "" {
		return false
	}
	for _, numOp := range numberOperators {
		if numOp == op {
			return true
		}
	}
	return false
}

// IsStringOperator returns true if the given operator can be used with a string value.
func IsStringOperator(op string) bool {
	op, _ = ParseOperator(op)
	if op == "" {
		return false
	}
	for _, strOp := range stringOperators {
		if strOp == op {
			return true
		}
	}
	return false
}

// ParseOperator returns the valid AQL operator for an arbitrary operator string.
// This supports different inputs, such as Filter.Operator in github.com/recipeer/go/qs
//
// If the input operator cannot be mapped to AQL, this function returns ErrInvalidOperator.
func ParseOperator(op string) (string, error) {
	if operators[op] == "" {
		return "", ErrInvalidOperator
	}
	return operators[op], nil
}
