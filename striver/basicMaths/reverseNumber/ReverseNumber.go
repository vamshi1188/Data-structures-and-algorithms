package main

import (
	"fmt"
	"math"
)

// Function to reverse the digits of an integer
func reverseNumber(n int) int {
	// TODO: Implement your solution here

	lastnum := 0
	reverseNumber := 0

	for n > 0 {
		lastnum = n % 10
		n = n / 10
		if reverseNumber > math.MaxInt32/10 || (reverseNumber == math.MaxInt32/10 && lastnum > 7) {
			return 0
		}
		if reverseNumber < math.MinInt32/10 || (reverseNumber == math.MinInt32/10 && lastnum < -8) {
			return 0
		}
		reverseNumber = reverseNumber*10 + lastnum

	}

	return reverseNumber // placeholder - replace with your solution
}

func main() {
	// Test your function here
	fmt.Println("Reverse Number Problem")
	fmt.Println("=====================")

	// Example test cases
	testCases := []int{12345, 7789, 10400, 123, 0, -123}

	for _, num := range testCases {
		result := reverseNumber(num)
		fmt.Printf("reverseNumber(%d) = %d\n", num, result)
	}

}
