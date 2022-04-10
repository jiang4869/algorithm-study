package main

func maximumCandies(candies []int, k int64) int {
	maxx := 0
	for _, val := range candies {
		maxx = max(maxx, val)
	}
	var l, r int64
	l, r = 1, int64(maxx)
	var res int
	for l <= r {
		mid := (l + r) / 2
		if check(candies, mid, k) {
			res = int(mid)
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return res
}
func check(candies []int, num, k int64) bool {
	if num == 0 {
		return true
	}
	sum := int64(0)
	for _, val := range candies {
		sum += int64(val) / num
	}

	return sum >= k
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*

0 8

*/
