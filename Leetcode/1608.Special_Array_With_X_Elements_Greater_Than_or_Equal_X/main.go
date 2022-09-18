package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	numsLen := len(nums)
	if nums[0] >= numsLen {
		return numsLen
	}
	for i := 1; i < numsLen; i++ {
		cnt := numsLen - i
		if nums[i] >= cnt && nums[i-1] < cnt {
			return cnt
		}
	}
	return -1
}
