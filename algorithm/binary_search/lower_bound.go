package main

import "sort"

// LowerBound  nums index for [begin,end)
func LowerBound(nums []int, begin, end, key int) int {
	l, r := begin, end
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= key {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func LowerBoundBySort(nums []int, key int) int {
	search := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= key
	})
	return search
}

func UpperBoundBySort(nums []int, key int) int {
	search := sort.Search(len(nums), func(i int) bool {
		return nums[i] > key
	})
	return search
}
