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

type node struct {
	X   int
	W   int
	Idx int
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(in, &n, &m)
		arr := make([]node, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &arr[i].X, &arr[i].W)
			arr[i].Idx = i + 1
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].W < arr[j].W
		})
		res := make([]node, 0)
		sum := 0
		for i := 0; i < n*2; i++ {
			res = append(res, arr[i])
			sum += arr[i].W
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i].X < res[j].X
		})
		fmt.Fprintln(out, sum)
		for i := 0; i < n; i++ {
			fmt.Fprintln(out, res[i].Idx, res[2*n-i-1].Idx)
		}

	}
}
