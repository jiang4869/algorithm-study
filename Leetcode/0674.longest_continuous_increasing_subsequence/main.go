package main

func findLengthOfLCIS(nums []int) int {
	//max := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	ans := 1
	begin := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			if ans < i-begin+1 {
				ans = i - begin + 1
			}
		} else {
			begin = i
		}
	}

	return ans
}
