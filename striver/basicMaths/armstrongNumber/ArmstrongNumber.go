package main

import (
	"fmt"
	"math"
)

// Function to check if a number is an Armstrong number
func isArmstrong(n int) bool {
	// TODO: Implement your solution here
	// Currently returns a placeholder value that will fail all tests
	// This ensures no tests pass until you implement the actual logic

	original := n
	count := 0
	temp := n

	// Step 1: Count number of digits
	for temp > 0 {
		count++
		temp /= 10
	}

	sum := 0
	temp = n

	// Step 2: Sum of each digit raised to the power of digit count
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(count)))
		temp /= 10
	}

	return sum == original
}

func main() {
	// // Test your function here
	fmt.Println("Armstrong Number Problem")
	fmt.Println("=======================")

	// // Example test cases
	testCases := []int{153, 371, 123, 9474, 1634, 0, 1, 407}

	for _, num := range testCases {
		result := isArmstrong(num)
		fmt.Printf("isArmstrong(%d) = %t\n", num, result)
	}
}
