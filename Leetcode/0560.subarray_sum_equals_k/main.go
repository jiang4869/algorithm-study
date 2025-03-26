package main

func subarraySum(nums []int, k int) int {
	ans := 0
	presum := make([]int, len(nums)+1)
	for i := range nums {
		presum[i+1] = presum[i] + nums[i]
	}
	mmid := make(map[int]int)

	for i := 0; i <= len(nums); i++ {
		num := presum[i] - k
		ans += mmid[num]
		mmid[presum[i]]++
	}
	return ans
}
