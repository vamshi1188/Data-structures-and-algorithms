package main

import "fmt"

func main() {

	num := 5
	pattern1(num)
	pattern2(num)
	pattern3(num)
	pattern4(num)
	pattern5(num)
}

// pattern functions

// *****
// ****
// ***
// **
// *

func pattern5(input int) {
	for i := 1; i <= input; i++ {
		for j := 0; j < input-i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 1
// 22
// 333
// 4444
// 55555

func pattern4(input int) {
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}

// 1
// 12
// 123
// 1234

func pattern3(input int) {
	for i := range input {
		for j := 1; j <= i+1; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

// *
// **
// ***
// ****

func pattern2(input int) {
	for i := range input {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// ****
// ****
// ****
// ****

func pattern1(input int) {
	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
