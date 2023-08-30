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
	var (
		t, x, y, n int
	)
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &x, &y, &n)
		tmp := n * (n - 1) / 2
		if x+tmp > y {
			fmt.Fprintln(out, "-1")
			continue
		}
		res := make([]int, 0)
		res = append(res, y)
		last := y
		delta := 1
		for i := 0; i < n-2; i++ {
			last = last - delta
			delta++
			res = append(res, last)
		}
		res = append(res, x)
		for i := n - 1; i >= 0; i-- {
			fmt.Fprint(out, res[i], " ")
		}
		fmt.Fprintln(out, "")
	}
}
