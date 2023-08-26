package main

func minimumSum(n int, k int) int {
	sum := 0
	i := 1
	for ; i <= n && i <= k/2; i++ {
		sum += i
	}
	tmp := k
	for ; i <= n; i++ {
		sum += tmp
		tmp++
	}
	return sum
}
