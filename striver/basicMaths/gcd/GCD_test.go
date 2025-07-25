package main

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		n1, n2   int
		expected int
	}{
		// Basic examples from problem statement
		{"Example 1: GCD(9, 12)", 9, 12, 3},
		{"Example 2: GCD(20, 15)", 20, 15, 5},

		// Additional test cases
		{"GCD(48, 18)", 48, 18, 6},
		{"GCD(7, 13) - coprime numbers", 7, 13, 1},
		{"GCD(100, 25)", 100, 25, 25},
		{"GCD(17, 19) - both prime", 17, 19, 1},
		{"GCD(56, 42)", 56, 42, 14},
		{"GCD(1071, 462)", 1071, 462, 21},

		// Edge cases
		{"GCD(0, 5)", 0, 5, 5},
		{"GCD(10, 0)", 10, 0, 10},
		{"GCD(1, 1)", 1, 1, 1},
		{"GCD(1, 100)", 1, 100, 1},

		// Negative numbers
		{"GCD(-12, 8)", -12, 8, 4},
		{"GCD(15, -25)", 15, -25, 5},
		{"GCD(-18, -24)", -18, -24, 6},

		// Same numbers
		{"GCD(25, 25)", 25, 25, 25},
		{"GCD(7, 7)", 7, 7, 7},

		// Large numbers
		{"GCD(1000, 500)", 1000, 500, 500},
		{"GCD(999, 333)", 999, 333, 333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gcd(tt.n1, tt.n2)
			if result != tt.expected {
				t.Errorf("gcd(%d, %d) = %d; expected %d", tt.n1, tt.n2, result, tt.expected)
			}
		})
	}
}

func TestGCDRecursive(t *testing.T) {
	tests := []struct {
		name     string
		n1, n2   int
		expected int
	}{
		// Basic examples from problem statement
		{"Example 1: GCD(9, 12)", 9, 12, 3},
		{"Example 2: GCD(20, 15)", 20, 15, 5},

		// Additional test cases
		{"GCD(48, 18)", 48, 18, 6},
		{"GCD(7, 13) - coprime numbers", 7, 13, 1},
		{"GCD(100, 25)", 100, 25, 25},
		{"GCD(17, 19) - both prime", 17, 19, 1},
		{"GCD(56, 42)", 56, 42, 14},
		{"GCD(1071, 462)", 1071, 462, 21},

		// Edge cases
		{"GCD(0, 5)", 0, 5, 5},
		{"GCD(10, 0)", 10, 0, 10},
		{"GCD(1, 1)", 1, 1, 1},
		{"GCD(1, 100)", 1, 100, 1},

		// Negative numbers
		{"GCD(-12, 8)", -12, 8, 4},
		{"GCD(15, -25)", 15, -25, 5},
		{"GCD(-18, -24)", -18, -24, 6},

		// Same numbers
		{"GCD(25, 25)", 25, 25, 25},
		{"GCD(7, 7)", 7, 7, 7},

		// Large numbers
		{"GCD(1000, 500)", 1000, 500, 500},
		{"GCD(999, 333)", 999, 333, 333},
	}

	for _, tt := range tests {
		t.Run(tt.name+" (Recursive)", func(t *testing.T) {
			result := gcdRecursive(tt.n1, tt.n2)
			if result != tt.expected {
				t.Errorf("gcdRecursive(%d, %d) = %d; expected %d", tt.n1, tt.n2, result, tt.expected)
			}
		})
	}
}

// Benchmark tests to compare iterative vs recursive performance
func BenchmarkGCDIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcd(1071, 462)
	}
}

func BenchmarkGCDRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcdRecursive(1071, 462)
	}
}
