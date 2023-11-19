package validate

// Validation error.
var (
	ErrValueNotAllowed = NewError("not allowed")
)

// In validates whether a value is found in a slice of allowed values.
func In[T comparable](allow ...T) func(T) error {
	return func(value T) error {
		for _, cmp := range allow {
			if cmp == value {
				return nil
			}
		}
		return ErrValueNotAllowed
	}
}

// NotIn validates whether a value is not found in a slice of disallowed values.
func NotIn[T comparable](allow ...T) func(T) error {
	return func(value T) error {
		for _, cmp := range allow {
			if cmp == value {
				return ErrValueNotAllowed
			}
		}
		return nil
	}
}
