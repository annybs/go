package validate

// Validation error.
var (
	ErrMustHaveMoreItems  = NewError("must have at least %d items")
	ErrMustHaveFewerItems = NewError("must have no more than %d items")
)

// MaxSize validates the length of a slice as being less than or equal to a given maximum.
func MaxSize[T any](l int) func([]T) error {
	return func(value []T) error {
		if len(value) > l {
			return ErrMustHaveFewerItems
		}
		return nil
	}
}

// MinSize validates the length of a slice as being greater than or equal to a given minimum.
func MinSize[T any](l int) func([]T) error {
	return func(value []T) error {
		if len(value) < l {
			return ErrMustHaveMoreItems
		}
		return nil
	}
}
