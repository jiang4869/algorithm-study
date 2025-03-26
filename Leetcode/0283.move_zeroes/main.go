package main

func moveZeroes(nums []int) {
	index := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
