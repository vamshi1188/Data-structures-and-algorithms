package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	ReverseArray(arr, 0, len(arr)-1)
	fmt.Println(arr)

}

func ReverseArray(array []int, l int, r int) {

	if l >= r {

		return
	}

	temp := array[l]
	array[l] = array[r]
	array[r] = temp

	ReverseArray(array, l+1, r-1)

}
