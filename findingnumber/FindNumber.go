package main

import "fmt"

func main() {
	numlist := []int{1, 2, 3, 4, 5}

	s := Findnum(numlist, 5)

	fmt.Println(s)

}

func Findnum(list []int, num int) bool {

	for _, i := range list {

		if i == num {
			return true
		}
	}

	return false
}
