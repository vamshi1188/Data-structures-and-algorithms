package main

import "fmt"

func main() {

	printName(1, 3)

}

func printName(i int, n int) {

	name := "vamshi"

	if i > n {
		return
	}
	fmt.Println(name)
	printName(i+1, n)
}
