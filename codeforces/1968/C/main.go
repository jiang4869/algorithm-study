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
		arr := make([]int, n-1)
		for i := 0; i < n-1; i++ {
			fmt.Fscan(in, &arr[i])
		}
		ans := make([]int, n)
		ans[1] = arr[0] + 50000
		ans[0] = ans[1] - arr[0]
		for i := 2; i < n; i++ {
			ans[i] = arr[i-1] + ans[i-1]
		}
		for i := 0; i < n; i++ {
			fmt.Fprint(out, ans[i], " ")
		}
		fmt.Fprintln(out, "")
	}
}

// 5 4 6 11
