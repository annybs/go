package arango

import (
	"errors"
)

// Operator error.
var (
	ErrInvalidOperator        = errors.New("invalid operator")
	ErrInvalidOperatorForType = errors.New("invalid operator for type")
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

	arrayOperators  = map[string]bool{"IN": true, "NOT IN": true}
	boolOperators   = map[string]bool{"==": true, "!=": true}
	numberOperators = map[string]bool{"==": true, "!=": true, ">": true, ">=": true, "<": true, "<=": true}
	stringOperators = map[string]bool{"==": true, "!=": true, ">": true, ">=": true, "<": true, "<=": true, "LIKE": true, "NOT LIKE": true}
)

// IsArrayOperator returns true if the given operator can be used with an array value.
func IsArrayOperator(op string) bool {
	_, err := ParseArrayOperator(op)
	return err == nil
}

// IsBoolOperator returns true if the given operator can be used with a Boolean value.
func IsBoolOperator(op string) bool {
	_, err := ParseBoolOperator(op)
	return err == nil
}

// IsNumberOperator returns true if the given operator can be used with a numeric value.
func IsNumberOperator(op string) bool {
	_, err := ParseNumberOperator(op)
	return err == nil
}

// IsStringOperator returns true if the given operator can be used with a string value.
func IsStringOperator(op string) bool {
	_, err := ParseStringOperator(op)
	return err == nil
}

// ParseArrayOperator returns the valid AQL operator for an array operator.
// It returns an error if the operator cannot be mapped to AQL or does not support arrays.
func ParseArrayOperator(op string) (string, error) {
	op, err := ParseOperator(op)
	if err != nil {
		return op, err
	}
	if !arrayOperators[op] {
		return op, ErrInvalidOperatorForType
	}
	return op, nil
}

// ParseBoolOperator returns the valid AQL operator for a Boolean operator.
// It returns an error if the operator cannot be mapped to AQL or does not support Booleans.
func ParseBoolOperator(op string) (string, error) {
	op, err := ParseOperator(op)
	if err != nil {
		return op, err
	}
	if !boolOperators[op] {
		return op, ErrInvalidOperatorForType
	}
	return op, nil
}

// ParseNumberOperator returns the valid AQL operator for a numeric operator.
// It returns an error if the operator cannot be mapped to AQL or does not support numbers.
func ParseNumberOperator(op string) (string, error) {
	op, err := ParseOperator(op)
	if err != nil {
		return op, err
	}
	if !numberOperators[op] {
		return op, ErrInvalidOperatorForType
	}
	return op, nil
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

// ParseStringOperator returns the valid AQL operator for a string operator.
// It returns an error if the operator cannot be mapped to AQL or does not support strings.
func ParseStringOperator(op string) (string, error) {
	op, err := ParseOperator(op)
	if err != nil {
		return op, err
	}
	if !stringOperators[op] {
		return op, ErrInvalidOperatorForType
	}
	return op, nil
}
