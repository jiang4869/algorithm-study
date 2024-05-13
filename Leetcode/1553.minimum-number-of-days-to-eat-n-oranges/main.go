package main

func minDays(n int) int {
	mmid := make(map[int]int)
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return i
		}
		if val, ok := mmid[i]; ok {
			return val
		}
		ans := min(dfs(i/3)+i%3, dfs(i/2)+i%2) + 1
		mmid[i] = ans
		return ans
	}
	return dfs(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
