package main

import (
	"testing"
)

func TestIsArmstrong(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		// Basic examples from problem statement
		{"Example 1: 153", 153, true},
		{"Example 2: 371", 371, true},

		// Single digits (all are Armstrong numbers)
		{"Single digit: 0", 0, true},
		{"Single digit: 1", 1, true},
		{"Single digit: 5", 5, true},
		{"Single digit: 9", 9, true},

		// Basic non-Armstrong numbers
		{"Not Armstrong: 123", 123, false},
		{"Not Armstrong: 10", 10, false},

		// More Armstrong numbers to ensure proper implementation
		{"Three digit Armstrong: 407", 407, true},
		{"Four digit Armstrong: 1634", 1634, true},

		// More non-Armstrong to ensure proper logic
		{"Not Armstrong: 111", 111, false},
		{"Not Armstrong: 1000", 1000, false},
	}

	// Check that the function is actually implemented (not just returning a constant)
	// Test a few known values to detect placeholder implementations
	if isArmstrong(153) == isArmstrong(123) && isArmstrong(371) == isArmstrong(10) {
		t.Fatal("Function appears to be returning a constant value. Please implement the actual Armstrong number logic.")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArmstrong(tt.input); got != tt.expected {
				if tt.expected {
					t.Errorf("isArmstrong(%d) = %t, want %t (this IS an Armstrong number)", tt.input, got, tt.expected)
				} else {
					t.Errorf("isArmstrong(%d) = %t, want %t (this is NOT an Armstrong number)", tt.input, got, tt.expected)
				}
			}
		})
	}
}
