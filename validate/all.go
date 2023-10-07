package validate

// All validates a value using a sequence of validation functions.
// If any validation function returns an error, the sequence stops and the error is returned.
func All[T any](fs ...func(T) error) func(T) error {
	return func(value T) error {
		for _, f := range fs {
			if err := f(value); err != nil {
				return err
			}
		}
		return nil
	}
}
