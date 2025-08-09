package main

import (
	"fmt"
)

func main() {

	printReverse(1, 5)

}
func printReverse(i, n int) {
	if i > n {
		return
	}
	fmt.Println(i)

	printReverse(i+1, n) // go deeper first
	fmt.Println(i)       // only runs on the way up
}
