package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	G := make([][]int, n+10)
	deps := make([]int, n+10)
	arr := make([]int, n+10)
	root := 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] == -1 {
			root = i
		} else {
			G[i] = append(G[i], arr[i])
			G[arr[i]] = append(G[arr[i]], i)
		}
	}
	maxDep := -1
	var dfs func(par, cur, dep int)
	sign := make([]bool, n+10)
	dfs = func(par, cur, dep int) {
		sign[cur] = true
		deps[cur] = dep
		for _, val := range G[cur] {
			if val != par && !sign[val] {
				dfs(cur, val, dep+1)
			}
		}
	}
	dfs(-1, root, 0)
	var num int
	mmid := make(map[int]int)
	mmid[root] = 1
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &num)
		if _, ok := mmid[num]; !ok {
			tmp := num
			for arr[num] != -1 {
				mmid[num] = 1
				num = arr[num]
				_, ok = mmid[num]
				if ok {
					break
				}
			}
			if deps[tmp] > maxDep {
				maxDep = deps[tmp]
			}
		}
		res := (len(mmid)-1)*2 - maxDep
		fmt.Fprintln(out, res)
	}
}
