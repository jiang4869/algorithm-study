package main

func removeDuplicates(nums []int) int {
	ans := 0
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		}
		if count < 3 {
			nums[ans] = nums[i]
			ans++
		} else {
			count = 0
		}
	}
	return ans
}
