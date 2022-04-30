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
	fmt.Fscan(in, &t, "\n")
	for ; t > 0; t-- {
		var l, r, a int
		fmt.Fscan(in, &l, &r, &a)
		var res int
		if a == 1 {
			res = r
		} else {
			p := (r/a)*a - 1
			if p < l {
				p = l
			}
			res = max(p/a+p%a, r/a+r%a)
		}
		fmt.Fprintln(out, res)

	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
