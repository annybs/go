package validate

import "errors"

var (
	ErrTooHigh = errors.New("too high")
	ErrTooLow  = errors.New("too low")
)

func Max(n int, exclusive bool) func(int) error {
	return func(value int) error {
		if exclusive {
			if value >= n {
				return ErrTooHigh
			}
		}
		if value > n {
			return ErrTooHigh
		}
		return nil
	}
}

func MaxFloat32(n float32, exclusive bool) func(float32) error {
	return func(value float32) error {
		if exclusive {
			if value >= n {
				return ErrTooHigh
			}
		}
		if value > n {
			return ErrTooHigh
		}
		return nil
	}
}

func MaxFloat64(n float64, exclusive bool) func(float64) error {
	return func(value float64) error {
		if exclusive {
			if value >= n {
				return ErrTooHigh
			}
		}
		if value > n {
			return ErrTooHigh
		}
		return nil
	}
}

func Min(n int, exclusive bool) func(int) error {
	return func(value int) error {
		if exclusive {
			if value <= n {
				return ErrTooHigh
			}
		}
		if value < n {
			return ErrTooHigh
		}
		return nil
	}
}

func MinFloat32(n float32, exclusive bool) func(float32) error {
	return func(value float32) error {
		if exclusive {
			if value <= n {
				return ErrTooHigh
			}
		}
		if value < n {
			return ErrTooHigh
		}
		return nil
	}
}

func MinFloat64(n float64, exclusive bool) func(float64) error {
	return func(value float64) error {
		if exclusive {
			if value <= n {
				return ErrTooHigh
			}
		}
		if value < n {
			return ErrTooHigh
		}
		return nil
	}
}
