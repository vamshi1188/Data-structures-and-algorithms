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
		{"Example 1: 153", 153, true},
		{"Example 2: 371", 371, true},
		{"Three digit Armstrong: 407", 407, true},
		{"Four digit Armstrong: 1634", 1634, true},
		{"Four digit Armstrong: 8208", 8208, true},
		{"Four digit Armstrong: 9474", 9474, true},
		{"Five digit Armstrong: 54748", 54748, true},
		{"Five digit Armstrong: 92727", 92727, true},
		{"Five digit Armstrong: 93084", 93084, true},
		{"Single digit: 0", 0, true},
		{"Single digit: 1", 1, true},
		{"Single digit: 5", 5, true},
		{"Single digit: 9", 9, true},
		{"Not Armstrong: 123", 123, false},
		{"Not Armstrong: 111", 111, false},
		{"Not Armstrong: 222", 222, false},
		{"Not Armstrong: 1000", 1000, false},
		{"Not Armstrong: 1635", 1635, false},
		{"Not Armstrong: 154", 154, false},
		{"Not Armstrong: 372", 372, false},
		{"Two digit: 10", 10, false},
		{"Two digit: 11", 11, false},
		{"Two digit: 99", 99, false},
		{"Large non-Armstrong: 12345", 12345, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArmstrong(tt.input); got != tt.expected {
				t.Errorf("isArmstrong(%d) = %t, want %t", tt.input, got, tt.expected)
			}
		})
	}
}

// Additional edge case tests
func TestIsArmstrongEdgeCases(t *testing.T) {
	edgeCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Zero", 0, true},
		{"Single digit 2", 2, true},
		{"Single digit 7", 7, true},
		{"Boundary: 100", 100, false},
		{"Boundary: 999", 999, false},
		{"Boundary: 1000", 1000, false},
		{"Six digit Armstrong: 548834", 548834, true},
		{"Large Armstrong: 1741725", 1741725, true},
		{"Large non-Armstrong: 1000000", 1000000, false},
	}

	for _, tt := range edgeCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArmstrong(tt.input); got != tt.expected {
				t.Errorf("isArmstrong(%d) = %t, want %t", tt.input, got, tt.expected)
			}
		})
	}
}
