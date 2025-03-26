package main

func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = numsLen + 1
		}
	}
	for _, num := range nums {
		num = abs(num)
		if 1 <= num && num <= numsLen {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}

	return numsLen + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
