package main

import "fmt"

func MissingNumber(nums []int) {

	number := make(map[int]bool)
	var missingnum int

	for i := range nums {

		if !number[nums[i]] {
			missingnum = i
		}
		if number[nums[i]] {

			number[nums[i]] = true
		} else {
			missingnum = i + 1
		}
	}

	fmt.Println(missingnum)

}
