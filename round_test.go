package maths_test

import (
	"fmt"
	"testing"

	"github.com/jabolopes/go-maths"
)

func ExampleDownToMultipleOf() {
	fmt.Printf("%d", maths.DownToMultipleOf(22, 10))
	// Output: 20
}

func ExampleUpToMultipleOf() {
	fmt.Printf("%d", maths.UpToMultipleOf(35, 10))
	// Output: 40
}

func TestDownToMultipleOf(t *testing.T) {
	tests := []struct{ n, m, want int }{
		{1, 1, 1},
		{2, 1, 2},
		{-11, 10, -20},
		{-9, 10, -10},
		{22, 10, 20},
		{35, 10, 30},
	}
	for _, test := range tests {
		got := maths.DownToMultipleOf(test.n, test.m)
		if got != test.want {
			t.Errorf("DownToMultipleOf(%v, %v) = %v, want %v", test.n, test.m, got, test.want)
		}
	}
}

func TestUpToMultipleOf(t *testing.T) {
	tests := []struct{ n, m, want int }{
		{1, 1, 1},
		{2, 1, 2},
		{-11, 10, -10},
		{-9, 10, 0},
		{22, 10, 30},
		{35, 10, 40},
	}
	for _, test := range tests {
		got := maths.UpToMultipleOf(test.n, test.m)
		if got != test.want {
			t.Errorf("UpToMultipleOf(%v, %v) = %v, want %v", test.n, test.m, got, test.want)
		}
	}
}
