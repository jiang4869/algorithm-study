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

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n+1)
		mmid := make(map[int]int)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &arr[i])
			mmid[arr[i]] = i
		}
		ans := make([]int, n+1)
		for i := n; i >= 1; i-- {

			ans[i] = (mmid[i]) % i
			for j := 1; j <= n; j++ {
				mmid[j] -= (mmid[i]) % i
				if mmid[j] < 1 {
					mmid[j] += i
				}
			}
		}
		for i := 1; i <= n; i++ {
			fmt.Fprintf(out, "%d ", ans[i])
		}
		fmt.Fprintln(out)
	}
}
