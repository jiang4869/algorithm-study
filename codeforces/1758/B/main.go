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
		var n int
		fmt.Fscan(in, &n)
		if n%2 == 0 {
			fmt.Fprintf(out, "1 3 ")
			for i := 0; i < n-2; i++ {
				fmt.Fprintf(out, "2 ")
			}
		} else {
			for i := 0; i < n; i++ {
				fmt.Fprintf(out, "1 ")
			}
		}
		fmt.Fprintln(out)
	}
}
