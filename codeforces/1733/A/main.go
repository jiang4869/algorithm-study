package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k int
		var num int64
		fmt.Fscan(in, &n, &k)
		res := make([]int64, k)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &num)
			res[i%k] = max(res[i%k], num)
		}
		sum := int64(0)
		for i := 0; i < k; i++ {
			sum += res[i]
		}
		fmt.Fprintln(out, sum)
	}
}
