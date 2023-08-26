package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximizeTheProfit(n int, offers [][]int) int {
	type pair struct{ start, gold int }
	dp := make([]int, n+1)
	f := make([][]pair, n)
	for i := 0; i < len(offers); i++ {
		f[offers[i][1]] = append(f[offers[i][1]], pair{start: offers[i][0], gold: offers[i][2]})

	}
	dp[0] = 0
	for end := 1; end <= n; end++ {
		dp[end] = dp[end-1]
		for _, v := range f[end-1] {
			dp[end] = max(dp[end], dp[v.start]+v.gold)
		}
	}
	return dp[n]
}
