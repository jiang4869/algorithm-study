package main

func minOperations(nums []int, target int) int {
	sum := int64(0)
	mmid := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sum += int64(nums[i])
		mmid[nums[i]]++
	}
	if sum < int64(target) {
		return -1
	}
	ans := 0
	sum = 0
	for i := 0; (1 << i) <= target; {
		sum += int64(mmid[1<<i]) << i
		mask := 1<<(i+1) - 1
		if sum >= int64((target & mask)) {
			i++
			continue
		}
		ans++
		for i++; mmid[1<<i] == 0; i++ {
			ans++
		}
	}
	return ans
}
