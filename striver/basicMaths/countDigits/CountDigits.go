package main

import "fmt"

// Function to count the number of digits in an integer
func countDigits(n int) int {
	// TODO: Implement your solution here

	count := 0

	numbers := n
	// lastDisgit := 1
	for numbers != 0 {
		numbers = numbers / 10
		count++
	}

	return count // placeholder - replace with your solution
}

func main() {
	// 	// Test your function here
	fmt.Println("Count Digits Problem")
	fmt.Println("===================")

	// 	// Example test cases
	testCases := []int{4, 14, 123, 0, -123, -4}

	for _, num := range testCases {
		result := countDigits(num)
		fmt.Printf("countDigits(%d) = %d\n", num, result)
	}

}
