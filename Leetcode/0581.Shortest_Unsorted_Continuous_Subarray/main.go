package main

import (
	"math"
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 0
	}
	sorted := sort.IsSorted(sort.IntSlice(nums))
	if sorted {
		return 0
	}

	var begin, end int
	begin, end = -1, -1
	max := math.MinInt32
	min := math.MaxInt32
	for i := 0; i < size; i++ {
		if max > nums[i] {
			end = i
		} else {
			max = nums[i]
		}
	}
	for i := size - 1; i >= 0; i-- {
		if min < nums[i] {
			begin = i
		} else {
			min = nums[i]
		}
	}

	return end - begin + 1
}
