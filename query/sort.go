package query

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

// Query error.
var (
	ErrInvalidSort  = errors.New("invalid sort")
	ErrTooManySorts = errors.New("too many sorts")
)

var sortRegexp = regexp.MustCompile("^([A-z0-9]+) (asc|desc)$")

// ReadSortsOptions configures the behaviour of ReadSorts.
type ReadSortsOptions struct {
	Key      string // Query string key. The default value is "sort"
	MaxSorts int    // If this is > 0, a maximum number of sorts is imposed
}

// Sort represents a sort order for, most likely, a database query.
type Sort struct {
	Field     string `json:"field"`     // Field by which to sort.
	Direction string `json:"direction"` // Direction in which to sort, namely asc or desc.
}

// ReadRequestSorts parses a request's query string into a slice of sorts.
// This function returns nil if no sorts are found.
func ReadRequestSorts(req *http.Request, opt *ReadSortsOptions) ([]Sort, error) {
	return ReadSorts(req.URL.Query(), opt)
}

// ReadSorts parses URL values into a slice of sorts.
// This function returns nil if no sorts are found.
func ReadSorts(values url.Values, opt *ReadSortsOptions) ([]Sort, error) {
	opt = initSortsOptions(opt)

	if !values.Has(opt.Key) {
		return nil, nil
	}

	if opt.MaxSorts > 0 && len(values[opt.Key]) > opt.MaxSorts {
		return nil, ErrTooManySorts
	}

	sorts := []Sort{}
	for _, sortStr := range values[opt.Key] {
		match := sortRegexp.FindStringSubmatch(sortStr)
		if match == nil {
			return nil, ErrInvalidSort
		}

		sort := Sort{
			Field:     match[1],
			Direction: match[2],
		}
		sorts = append(sorts, sort)
	}

	return sorts, nil
}

// ReadStringSorts parses a query string literal into a slice of sorts.
// This function returns nil if no sorts are found.
func ReadStringSorts(qs string, opt *ReadSortsOptions) ([]Sort, error) {
	values, err := url.ParseQuery(qs)
	if err != nil {
		return nil, err
	}
	return ReadSorts(values, opt)
}

func initSortsOptions(opt *ReadSortsOptions) *ReadSortsOptions {
	def := &ReadSortsOptions{
		Key: "sort",
	}

	if opt != nil {
		if len(opt.Key) > 0 {
			def.Key = opt.Key
		}

		if opt.MaxSorts > def.MaxSorts {
			def.MaxSorts = opt.MaxSorts
		}
	}

	return def
}
