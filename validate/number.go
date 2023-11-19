package validate

// Validation error.
var (
	ErrMustBeGreater        = NewError("must be greater than %v")
	ErrMustBeGreaterOrEqual = NewError("must be greater than or equal to %v")
	ErrMustBeLess           = NewError("must be less than %v")
	ErrMustBeLessOrEqual    = NewError("must be less than or equal to %v")
)

// Max validates whether an integer is less than or equal to a given maximum.
// If exclusive is true, an equal value will also produce an error.
func Max(n int, exclusive bool) func(int) error {
	return func(value int) error {
		if exclusive {
			if value >= n {
				return ErrMustBeLess.With(n)
			}
		}
		if value > n {
			return ErrMustBeLessOrEqual.With(n)
		}
		return nil
	}
}

// MaxFloat32 validates whether a float32 is less than or equal to a given maximum.
// If exclusive is true, an equal value will also produce an error.
func MaxFloat32(n float32, exclusive bool) func(float32) error {
	return func(value float32) error {
		if exclusive {
			if value >= n {
				return ErrMustBeLess.With(n)
			}
		}
		if value > n {
			return ErrMustBeLessOrEqual.With(n)
		}
		return nil
	}
}

// MaxFloat64 validates whether a float64 is less than or equal to a given maximum.
// If exclusive is true, an equal value will also produce an error.
func MaxFloat64(n float64, exclusive bool) func(float64) error {
	return func(value float64) error {
		if exclusive {
			if value >= n {
				return ErrMustBeLess.With(n)
			}
		}
		if value > n {
			return ErrMustBeLessOrEqual.With(n)
		}
		return nil
	}
}

// Min validates whether an integer is less than or equal to a given maximum.
// If exclusive is true, an equal value will also produce an error.
func Min(n int, exclusive bool) func(int) error {
	return func(value int) error {
		if exclusive {
			if value <= n {
				return ErrMustBeGreater.With(n)
			}
		}
		if value < n {
			return ErrMustBeGreaterOrEqual.With(n)
		}
		return nil
	}
}

// MinFloat32 validates whether a float32 is less than or equal to a given maximum.
// If exclusive is true, an equal value will also produce an error.
func MinFloat32(n float32, exclusive bool) func(float32) error {
	return func(value float32) error {
		if exclusive {
			if value <= n {
				return ErrMustBeGreater.With(n)
			}
		}
		if value < n {
			return ErrMustBeGreaterOrEqual.With(n)
		}
		return nil
	}
}

// MinFloat64 validates whether a float64 is less than or equal to a given maximum.
// If exclusive is true, an equal value will also produce an error.
func MinFloat64(n float64, exclusive bool) func(float64) error {
	return func(value float64) error {
		if exclusive {
			if value <= n {
				return ErrMustBeGreater.With(n)
			}
		}
		if value < n {
			return ErrMustBeGreaterOrEqual.With(n)
		}
		return nil
	}
}
