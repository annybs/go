package arango

import (
	"fmt"
	"strings"
)

// Query provides a simple way to build ArangoDB queries.
type Query struct {
	Lines  []string
	Params map[string]any
}

// Append adds a line to the query, attaching (optional) values to bind parameters.
//
// Any variables given are assigned to bind parameters in the same order as they appear in the line.
// Keep in mind the following behaviours:
//
//   - If more variables are provided than there are parameters in the line, leftover variables are discarded
//   - If more parameters are used than variables are given, remaining parameters are left unmapped
//   - Reused parameters are overwritten
func (query *Query) Append(line string, values ...any) *Query {
	var params map[string]any = nil

	names := ReadParams(line)
	if len(names) > 0 {
		params = map[string]any{}
		for i, name := range names {
			if i == len(values) {
				break
			}
			params[name] = values[i]
		}
	}

	query.Lines = append(query.Lines, line)
	return query.BindMap(params)
}

// Appendf adds a line to the query, replacing values directly into the string using standard formatting syntax.
// This is a shorthand for:
//
//	// query *Query
//	query.Append(fmt.Sprintf("my format (%s)", "example"))
//
// This method does not attach values to bind parameters.
func (query *Query) Appendf(line string, a ...any) *Query {
	line = fmt.Sprintf(line, a...)
	return query.Append(line)
}

// Bind binds a value to a bind parameter.
func (query *Query) Bind(name string, value any) *Query {
	query.Params[name] = value
	return query
}

// BindMap assigns values to bind parameters.
func (query *Query) BindMap(params map[string]any) *Query {
	if params != nil {
		for name, value := range params {
			query.Params[name] = value
		}
	}
	return query
}

// Copy creates a copy of the query.
func (query *Query) Copy() *Query {
	newQuery := NewQuery()
	for _, line := range query.Lines {
		newQuery.Lines = append(newQuery.Lines, line)
	}
	for name, value := range query.Params {
		newQuery.Params[name] = value
	}
	return newQuery
}

func (query *Query) String() string {
	return strings.Join(query.Lines, "\n")
}

// NewQuery creates a new, empty Query.
func NewQuery() *Query {
	return &Query{
		Lines:  []string{},
		Params: map[string]any{},
	}
}
