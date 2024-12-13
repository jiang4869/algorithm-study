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
		for i := 1; i <= n-2; i++ {
			fmt.Fprintln(out, i, i)
		}
		fmt.Fprintln(out, n-1, n)
		fmt.Fprintln(out, n, n)
	}
}
