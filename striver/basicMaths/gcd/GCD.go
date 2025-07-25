package main

import "fmt"

// Function to find the Greatest Common Divisor (GCD) of two integers
func gcd(n1, n2 int) int {
	// TODO: Implement your solution here
	// Currently returns a placeholder value that will fail most tests
	// This ensures no tests pass until you implement the actual logic

	// Handle edge cases
	if n1 == 0 {
		return n2
	}
	if n2 == 0 {
		return n1
	}

	// Make both numbers positive for calculation
	if n1 < 0 {
		n1 = -n1
	}
	if n2 < 0 {
		n2 = -n2
	}

	// Euclidean Algorithm - Iterative approach
	for n2 != 0 {
		temp := n2
		n2 = n1 % n2
		n1 = temp
	}

	return n1 // placeholder - replace with your solution
}

// Alternative recursive implementation using Euclidean algorithm
func gcdRecursive(n1, n2 int) int {
	// Handle edge cases
	if n1 == 0 {
		return n2
	}
	if n2 == 0 {
		return n1
	}

	// Make both numbers positive for calculation
	if n1 < 0 {
		n1 = -n1
	}
	if n2 < 0 {
		n2 = -n2
	}

	// Recursive Euclidean Algorithm
	if n2 == 0 {
		return n1
	}
	return gcdRecursive(n2, n1%n2)
}

func main() {
	// Test your function here
	fmt.Println("Greatest Common Divisor (GCD) Problem")
	fmt.Println("====================================")

	// Example test cases from problem statement
	testCases := []struct {
		n1, n2   int
		expected int
	}{
		{9, 12, 3},
		{20, 15, 5},
		{48, 18, 6},
		{7, 13, 1},
		{0, 5, 5},
		{10, 0, 10},
		{-12, 8, 4},
		{15, -25, 5},
		{1, 1, 1},
		{100, 25, 25},
	}

	for _, tc := range testCases {
		result := gcd(tc.n1, tc.n2)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("%s GCD(%d, %d) = %d (expected: %d)\n", 
			status, tc.n1, tc.n2, result, tc.expected)
	}

	fmt.Println("\n--- Testing Recursive Implementation ---")
	for _, tc := range testCases {
		result := gcdRecursive(tc.n1, tc.n2)
		status := "✓"
		if result != tc.expected {
			status = "✗"
		}
		fmt.Printf("%s GCD_Recursive(%d, %d) = %d (expected: %d)\n", 
			status, tc.n1, tc.n2, result, tc.expected)
	}
}
