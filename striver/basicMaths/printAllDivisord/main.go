package main

import (
	"fmt"
	"math"
)

func main() {
	num := 27
	fmt.Println(PrintDivisors(num)) // true
}

func PrintDivisors(num int) bool {
	if num <= 1 {
		return false // 1 or less can't be perfect
	}

	sum := 1 // Start with 1, always a proper divisor

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			pair := num / i

			if i != pair {
				sum += i + pair
			} else {
				sum += i // add square root only once
			}
		}
	}

	return sum == num
}
