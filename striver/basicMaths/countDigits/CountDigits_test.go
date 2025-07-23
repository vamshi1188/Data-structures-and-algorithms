package main

import (
	"testing"
)

func TestCountDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Example 1: Single digit", 4, 1},
		{"Example 2: Two digits", 14, 2},
		{"Three digit positive", 123, 3},
		{"Four digit with zeros", 1000, 4},
		{"Zero special case", 0, 1},
		{"Single digit negative", -4, 1},
		{"Three digit negative", -123, 3},
		{"Large positive number", 99999, 5},
		{"Maximum int32", 2147483647, 10},
		{"Minimum int32", -2147483648, 10},
		{"Single digit 9", 9, 1},
		{"Two digit 10", 10, 2},
		{"Two digit 99", 99, 2},
		{"Three digit 100", 100, 3},
		{"Six digits", 123456, 6},
		{"Seven digits", 1234567, 7},
		{"Eight digits", 12345678, 8},
		{"Nine digits", 123456789, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.input); got != tt.expected {
				t.Errorf("countDigits(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}

// Additional edge case tests
func TestCountDigitsEdgeCases(t *testing.T) {
	edgeCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"Negative single digit", -1, 1},
		{"Negative two digits", -12, 2},
		{"Negative large number", -987654321, 9},
		{"Positive single digit boundary", 1, 1},
		{"Positive single digit max", 9, 1},
		{"Positive two digit min", 10, 2},
		{"Positive two digit max", 99, 2},
		{"Positive three digit min", 100, 3},
		{"Positive three digit max", 999, 3},
	}

	for _, tt := range edgeCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.input); got != tt.expected {
				t.Errorf("countDigits(%d) = %d, want %d", tt.input, got, tt.expected)
			}
		})
	}
}
