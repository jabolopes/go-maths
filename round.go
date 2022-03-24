package maths

import "golang.org/x/exp/constraints"

// Down rounds n down to nearest multiple of m
func DownToMultipleOf[T constraints.Integer](n, m T) T {
	if n >= 0 {
		return (n / m) * m
	}
	return ((n - m + 1) / m) * m
}

// Up rounds n up to nearest multiple of m
//
// This is actually -1 * Down(-n, m)
func UpToMultipleOf[T constraints.Integer](n, m T) T {
	if n >= 0 {
		return ((n + m - 1) / m) * m
	}

	return (n / m) * m
}
