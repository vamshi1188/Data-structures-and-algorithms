package main

import (
	"fmt"
)

func main() {

	num := 4
	pattern1(num)
	pattern2(num)
	pattern3(num)
	pattern4(num)
	pattern5(num)
	pattern6(num)
	pattern7(num)
	pattern8(num)
	pattern9(num)
	pattern10(num)
	pattern11(num)
	pattern12(num)
	pattern13(num)
	pattern14(num)
	pattern15(num)
	pattern16(num)
	pattern17(num)
	pattern18(num)
	pattern19(num)
	pattern20(num)
	pattern21(num)
	pattern22(num)
}

// pattern22

func pattern22(num int) {
	fmt.Println("pattern 22")

	for i := 0; i < 2*num-1; i++ {

		for j := 0; j < 2*num-1; j++ {

			top := i
			left := j
			right := (2*num - 2) - j
			down := (2*num - 2) - i
			fmt.Print(num-min(min(top, down), min(left, right)), " ")
		}
		fmt.Println()
	}
}

// pattern21

func pattern21(num int) {
	fmt.Println("pattern 21")

	for i := 0; i < num; i++ {

		for j := 0; j < num; j++ {
			if j == 0 || i == 0 || j == num-1 || i == num-1 {

				fmt.Print("* ")

			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}
}

// patter20
// *        *
// **      **
// ***    ***
// ****  ****
// **********
// **********
// ****  ****
// ***    ***
// **      **
// *        *

func pattern20(num int) {
	fmt.Println("pattern20")

	spacex := 8
	for i := 1; i <= num; i++ {

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		// spaces
		for j := 0; j < spacex; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {

			fmt.Print("*")
		}
		fmt.Println()
		spacex -= 2
	}
	space := 0

	for i := 1; i <= num; i++ {

		for j := 0; j <= num-i; j++ {

			fmt.Print("*")
		}

		for j := 0; j < space; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= num-i; j++ {
			fmt.Print("*")
		}
		space += 2

		fmt.Println()

	}
}

// pattern19

func pattern19(num int) {

	space := 0
	fmt.Println("pattern 12")

	for i := 1; i <= num; i++ {

		for j := 0; j <= num-i; j++ {

			fmt.Print("*")
		}

		for j := 0; j < space; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= num-i; j++ {
			fmt.Print("*")
		}
		space += 2

		fmt.Println()

	}

	spacex := 8
	for i := 1; i <= num; i++ {

		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		// spaces
		for j := 0; j < spacex; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < i; j++ {

			fmt.Print("*")
		}
		fmt.Println()
		spacex -= 2
	}
}

// pattern18

func pattern18(num int) {
	for i := 0; i < num; i++ {

		for j := 'E' - rune(i); j <= 'E'; j++ {

			fmt.Printf("%c", j)

		}

		fmt.Println()
	}
}

// pattern17

func pattern17(num int) {
	for i := 0; i < num; i++ {

		for j := 0; j < num-i-1; j++ {
			fmt.Print(" ")
		}

		// alphabets
		ch := 'A'
		breakpoint := (2*i + 1) / 2
		for j := 1; j <= 2*i+1; j++ {
			fmt.Printf("%c", ch)
			if j <= breakpoint {
				ch++
			} else {
				ch--
			}
		}

		for k := 0; k < num-i-1; k++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// pattern16

func pattern16(num int) {
	fmt.Println("pattern 16")

	for i := 0; i < num; i++ {
		char := 'A' + i
		for j := 'A'; j <= rune(char); j++ {

			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
}

// pattern15

func pattern15(num int) {
	fmt.Println("pattern 15")
	for i := 0; i < num; i++ {
		for ch := 'A'; ch <= 'A'+rune(num-i-1); ch++ {

			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

// pattern14

func pattern14(num int) {
	for i := 0; i < num; i++ {
		for ch := 'A'; ch <= 'A'+rune(i); ch++ {

			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

// pattern 13

func pattern13(num int) {

	number := 1

	for i := 1; i <= num; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(number)
			number += 1
		}
		fmt.Println()
	}
}

// pattern 12

func pattern12(num int) {
	space := 2 * (num - 2)
	fmt.Println("pattern 12")

	for i := 1; i < num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		// spaces
		for j := 1; j <= space; j++ {
			fmt.Print(" ")
		}

		for j := i; j >= 1; j-- {

			fmt.Print(j)
		}
		fmt.Println()
		space -= 2

	}

}

// patterm 11

func pattern11(num int) {

	var start = 0
	for i := 0; i < num; i++ {
		if i%2 == 0 {
			start = 0
		} else {
			start = 1
		}
		for j := 0; j < i; j++ {

			fmt.Print(start)

			start = 1 - start
		}
		fmt.Println()
	}
}

// pattern 10

func pattern10(num int) {
	fmt.Println("pattern 10")
	for i := 0; i < num; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < num; i++ {
		for j := 1; j < num-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// pattern 9
//     *
//    ***
//   *****
//  *******
// *********
// *********
//  *******
//   *****
//    ***
//     *

func pattern9(num int) {

	fmt.Println("pattern 9")

	for i := 0; i < num; i++ {

		for j := 0; j < num-i-1; j++ {
			fmt.Print(" ")
		}

		for h := 0; h < 2*i+1; h++ {

			fmt.Print("*")
		}
		for k := 0; k < num-i-1; k++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
	for i := num - 1; i >= 0; i-- {

		for j := 0; j < num-i-1; j++ {
			fmt.Print(" ")
		}

		for h := 0; h < 2*i+1; h++ {

			fmt.Print("*")
		}

		fmt.Println()
	}
}

// pattern 8
// *********
//  *******
//   *****
//    ***
//     *

func pattern8(num int) {
	fmt.Println("pattern 8")

	for i := num; i >= 0; i-- {

		for j := 0; j < num-i; j++ {
			fmt.Print(" ")
		}

		for h := 0; h < 2*i+1; h++ {

			fmt.Print("*")
		}

		fmt.Println()
	}
}

// pattern
//     *
//    ***
//   *****
//  *******
// *********

func pattern7(num int) {
	for i := 0; i < num; i++ {

		for j := 0; j < num-i-1; j++ {
			fmt.Print(" ")
		}

		for h := 0; h < 2*i+1; h++ {

			fmt.Print("*")
		}
		for k := 0; k < num-i-1; k++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

//  pattern
// 1234
// 123
// 12
// 1

func pattern6(input int) {

	for i := 1; i <= input; i++ {

		for j := 1; j < input-i+1; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
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
