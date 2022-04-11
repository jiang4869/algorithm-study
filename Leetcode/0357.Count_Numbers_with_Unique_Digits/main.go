package main

func countNumbersWithUniqueDigits(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		tmp := 9
		for j := 0; j < i; j++ {
			tmp *= 9 - j
		}
		res += tmp
	}
	return res
}
