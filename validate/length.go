package validate

// Validation error.
var (
	ErrMustBeLonger  = NewError("must be at least %d characters")
	ErrMustBeShorter = NewError("must be no more than %d characters")
)

// MaxLength validates the length of a string as being less than or equal to a given maximum.
func MaxLength(l int) func(string) error {
	return func(value string) error {
		if len(value) > l {
			return ErrMustBeShorter.With(l)
		}
		return nil
	}
}

// MinLength validates the length of a string as being greater than or equal to a given minimum.
func MinLength(l int) func(string) error {
	return func(value string) error {
		if len(value) < l {
			return ErrMustBeLonger.With(l)
		}
		return nil
	}
}
