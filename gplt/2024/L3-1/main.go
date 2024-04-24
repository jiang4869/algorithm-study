package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	sort.Ints(arr)
	res := make([]int, 0)
	var dfs func(cur, sum int, s []int)
	flag := false
	dfs = func(cur, sum int, s []int) {
		if flag {
			return
		}
		if sum > m {
			return
		}
		if sum == m {
			flag = true
			return
		}
		dfs(cur+1, sum, s)
		s = append(s, arr[cur])
		dfs(cur+1, sum+arr[cur], s)
		s = s[:len(s)-1]
	}
	dfs(0, 0, []int{})
	if flag {
		for i := 0; i < len(res); i++ {
			if i != 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, res[i])
		}
	} else {
		fmt.Fprintln(out, "No Solution")
	}
}
