package maths_test

import (
	"fmt"
	"testing"

	"github.com/jabolopes/go-maths"
)

func ExampleMath() {
	fmt.Printf("%d\n", maths.Abs(-1))
	fmt.Printf("%d\n", maths.Neg(-2))
	fmt.Printf("%d\n", maths.Min(1, 2))
	fmt.Printf("%d\n", maths.Max(1, 2))
	fmt.Printf("%d\n", maths.Clamp(11, 1, 10))
	// Output: 1
	// -2
	// 1
	// 2
	// 10
}

func TestAbs(t *testing.T) {
	tests := []struct{ value, want int }{
		{0, 0},
		{-1, 1},
		{1, 1},
	}
	for _, test := range tests {
		got := maths.Abs(test.value)
		if got != test.want {
			t.Errorf("Abs(%v) = %v, want %v", test.value, got, test.want)
		}
	}
}

func TestNeg(t *testing.T) {
	tests := []struct{ value, want int }{
		{0, 0},
		{-1, -1},
		{1, -1},
	}
	for _, test := range tests {
		got := maths.Neg(test.value)
		if got != test.want {
			t.Errorf("Neg(%v) = %v, want %v", test.value, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{0, 0, 0},
		{0, 1, 0},
		{-1, 1, -1},
		{1, 0, 0},
		{1, -1, -1},
	}
	for _, test := range tests {
		got := maths.Min(test.a, test.b)
		if got != test.want {
			t.Errorf("Min(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{0, 0, 0},
		{0, 1, 1},
		{-1, 1, 1},
		{1, 0, 1},
		{1, -1, 1},
	}
	for _, test := range tests {
		got := maths.Max(test.a, test.b)
		if got != test.want {
			t.Errorf("Max(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct{ value, min, max, want int }{
		// Min.
		{-1, 0, 2, 0},
		{0, 0, 2, 0},
		// Mid.
		{1, 0, 2, 1},
		{2, 0, 2, 2},
		// Max.
		{2, 0, 2, 2},
		{3, 0, 2, 2},
		// Empty range.
		{-1, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 0},
	}
	for _, test := range tests {
		got := maths.Clamp(test.value, test.min, test.max)
		if got != test.want {
			t.Errorf("Clamp(%v, %v, %v) = %v, want %v", test.value, test.min, test.max, got, test.want)
		}
	}
}
