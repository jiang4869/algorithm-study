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
		t   int
		n   int
		num int
	)
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &n)
		count := 0

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &num)
			if num%2 == 1 {
				count++
			}
		}
		if count%2 == 0 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
