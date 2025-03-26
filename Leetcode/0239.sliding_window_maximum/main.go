package main

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0, len(nums))
	deque := make([]int, 0)
	for i, num := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if deque[0]+k-1 < i {
			deque = deque[0:]
		}
		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}
	return ans
}
