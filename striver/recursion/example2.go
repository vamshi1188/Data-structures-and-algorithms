package main

import "fmt"

func main() {

	PrintNum(1, 5)
}

func PrintNum(i int, n int) {
	if i > n {
		return
	}

	PrintNum(i+1, n)

	fmt.Println(i)
}
