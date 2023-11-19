package validate

var (
	ErrNotEqual = NewError("must be equal to %v")
)

// Equal validates whether an input value is equal to a comparison value.
func Equal[T comparable](cmp T) func(T) error {
	return func(value T) error {
		if value != cmp {
			return ErrNotEqual.With(cmp)
		}
		return nil
	}
}
