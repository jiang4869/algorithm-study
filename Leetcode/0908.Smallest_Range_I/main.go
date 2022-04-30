package main

func smallestRangeI(nums []int, k int) int {
	minn, maxn := nums[0], nums[0]
	for _, val := range nums {
		minn = min(minn, val)
		maxn = max(maxn, val)
	}

	return max(0, maxn-minn-2*k)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
