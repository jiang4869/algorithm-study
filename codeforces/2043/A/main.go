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
		var n int64
		fmt.Fscan(in, &n)
		res := int64(1)
		for n >= 4 {
			num := n / 4
			res *= 2
			n = num
		}
		fmt.Fprintln(out, res)
	}
}
