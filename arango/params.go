package arango

import (
	"regexp"
)

// https://docs.arangodb.com/3.11/aql/fundamentals/bind-parameters/#syntax
var paramRegexp = regexp.MustCompile("@(@?[A-z0-9_]+)")

// ReadParams reads out named parameters from an AQL string.
func ReadParams(input string) []string {
	params := []string{}

	matches := paramRegexp.FindAllStringSubmatch(input, -1)
	if matches != nil {
		for _, match := range matches {
			params = append(params, match[1])
		}
	}

	return params
}
