package main

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	ans := 0
	arr := make([]int, 1e5+10)
	for i := 0; i < len(difficulty); i++ {
		if arr[difficulty[i]] < profit[i] {
			arr[difficulty[i]] = profit[i]
		}
	}
	for i := 1; i <= 1e5; i++ {
		if arr[i] < arr[i-1] {
			arr[i] = arr[i-1]
		}
	}
	for i := 0; i < len(worker); i++ {
		ans += arr[worker[i]]
	}
	return ans
}
