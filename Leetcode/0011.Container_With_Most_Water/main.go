package main

func maxArea(height []int) int {
	ans := 0
	length := len(height)
	l, r := 0, length-1
	for l < r {
		ans = Max(ans, Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
