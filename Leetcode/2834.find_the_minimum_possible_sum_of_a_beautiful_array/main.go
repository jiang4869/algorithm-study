package main

func minimumPossibleSum(n int, target int) int64 {
	sum := int64(0)
	i := 1
	for ; i <= n && i <= target/2; i++ {
		sum += int64(i)
	}
	tmp := target
	for ; i <= n; i++ {
		sum += int64(tmp)
		tmp++
	}
	return sum
}
