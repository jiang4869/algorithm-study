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
		var n, x, y int
		fmt.Fscan(in, &n, &x, &y)
		if (x+y) == 0 || ((x != 0 && y != 0) || (n-1)%(x+y) != 0) {
			fmt.Fprintln(out, "-1")
		} else {
			num := 1
			flag := x
			cur := 0
			for i := 1; i < n; i++ {
				if cur < flag {
					cur++
				} else {
					num = i + 1
					cur = 1
					if flag == x {
						flag = y
					} else {
						flag = x
					}
					if flag == 0 {
						if flag == x {
							flag = y
						} else {
							flag = x
						}
					}
				}
				fmt.Fprintf(out, "%d ", num)
			}
			fmt.Fprintln(out, "")
		}
	}
}
