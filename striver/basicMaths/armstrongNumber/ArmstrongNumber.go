package main

import "fmt"

// Function to check if a number is an Armstrong number
func isArmstrong(n int) bool {
	// TODO: Implement your solution here
	return false // placeholder - replace with your solution
}

func main() {
	// Test your function here
	fmt.Println("Armstrong Number Problem")
	fmt.Println("=======================")
	
	// Example test cases
	testCases := []int{153, 371, 123, 9474, 1634, 0, 1, 407}
	
	for _, num := range testCases {
		result := isArmstrong(num)
		fmt.Printf("isArmstrong(%d) = %t\n", num, result)
	}
}
