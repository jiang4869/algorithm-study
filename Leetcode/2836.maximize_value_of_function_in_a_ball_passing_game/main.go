package main

import "math/bits"

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func getMaxFunctionValue(receiver []int, K int64) int64 {
	type pair struct{ pa, sum int }
	n, m := len(receiver), bits.Len(uint(K))
	pa := make([][]pair, n)
	for i, p := range receiver {
		pa[i] = make([]pair, m)
		pa[i][0] = pair{p, p}
	}

	for i := 0; i+1 < m; i++ {
		for x := range pa {
			p := pa[x][i]
			pp := pa[p.pa][i]
			pa[x][i+1] = pair{pp.pa, pp.sum + p.sum}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		x := i
		sum := i
		for k := 0; k-1 < m; k++ {
			if ((K >> k) & 1) > 0 {
				p := pa[x][k]
				sum += p.sum
				x = p.pa
			}
		}
		ans = max(ans, sum)
	}
	return int64(ans)
}
