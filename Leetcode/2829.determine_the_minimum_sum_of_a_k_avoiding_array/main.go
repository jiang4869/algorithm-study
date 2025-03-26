package main

func minimumSum(n int, k int) int {
	ans := 0
	a := k / 2
	if a >= n {
		ans = (1 + a) * a / 2
	} else {
		ans = (1 + a) * a / 2
		ans += (k + (n - a) - 1) * (n - 1) / 2
	}

	return ans

}
