package main

import "sort"

// O(n)解法
func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	ans := 0
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]+nums[r] < target {
			ans += r - l
			l++
		} else {
			r--
		}
	}
	return ans
}

// 暴力解法
//func countPairs(nums []int, target int) int {
//	ans := 0
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] < target {
//				ans++
//			}
//		}
//	}
//	return ans
//}
