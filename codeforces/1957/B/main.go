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
		var n, k int
		fmt.Fscan(in, &n, &k)
		if n == 1 {
			fmt.Fprintln(out, k)
			continue
		}
		msb := 0
		for i := 0; i < 31; i++ {
			if (k & (1 << i)) != 0 {
				msb = i
			}
		}
		arr := make([]int, n)
		arr[0] = (1 << msb) - 1
		arr[1] = k - arr[0]
		for i := 0; i < n; i++ {
			fmt.Fprint(out, arr[i], " ")
		}
		fmt.Fprintln(out, "")
	}
}
