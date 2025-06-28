package main

func FIndDuplicateNum(nums []int) int {
	var dupliacte int
	temp := make(map[int]bool)

	for i := 0; i < len(nums); i++ {

		if temp[nums[i]] {

			dupliacte = nums[i]
		} else {
			temp[nums[i]] = true
		}
	}
	return dupliacte
}
