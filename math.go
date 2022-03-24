package maths

import "golang.org/x/exp/constraints"

type number interface {
	constraints.Integer | constraints.Float
}

// Abs returns the absolute of value.
func Abs[T number](value T) T {
	if value < 0 {
		return -value
	}

	return value
}

// Neg returns the negative of value. It does not negate value. For
// negating, simply use -value instead.
func Neg[T number](value T) T {
	if value < 0 {
		return value
	}

	return -value
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Clamp clamps value to the range [min, max]. Returns value if value
// is already in the range [min, max]. Otherwise, it either returns
// min or max depending on whether value is less than min or greater
// than max, respectively.
func Clamp[T constraints.Ordered](value, min, max T) T {
	if min > max {
		return min
	}

	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}
