package main

import (
	"testing"
)

func TestReverseNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Example 1: Basic reverse", 12345, 54321},
		{"Example 2: Basic reverse", 7789, 9877},
		{"Trailing zeros removed", 10400, 401},
		{"Single digit", 5, 5},
		{"Zero", 0, 0},
		{"Two digits", 12, 21},
		{"Three digits", 123, 321},
		{"Number ending with zero", 120, 21},
		{"Number ending with multiple zeros", 1000, 1},
		{"Large number", 123456789, 987654321},
		{"Negative single digit", -5, -5},
		{"Negative two digits", -12, -21},
		{"Negative three digits", -123, -321},
		{"Negative with trailing zeros", -120, -21},
		{"Negative large number", -12345, -54321},
		{"Maximum single digit", 9, 9},
		{"Minimum two digit", 10, 1},
		{"Maximum two digit", 99, 99},
		{"Palindrome", 121, 121},
		{"Another palindrome", 1221, 1221},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseNumber(tt.input); got != tt.expected {
				t.Errorf("reverseNumber(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

// Additional edge case tests
func TestReverseNumberEdgeCases(t *testing.T) {
	edgeCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Multiple trailing zeros", 100200, 201},
		{"All zeros except first", 100000, 1},
		{"Negative with multiple trailing zeros", -100200, -201},
		{"Large negative", -987654321, -123456789},
		{"Single zero", 0, 0},
		{"Negative zero equivalent", -0, 0},
		{"Powers of 10", 1000000, 1},
		{"Negative powers of 10", -1000000, -1},
	}

	for _, tt := range edgeCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseNumber(tt.input); got != tt.expected {
				t.Errorf("reverseNumber(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
