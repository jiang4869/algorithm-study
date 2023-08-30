package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
		res := make([]int, 0)
		res = append(res, n)
		for n > 1 {
			x := bits.TrailingZeros(uint(n))
			if (1 << x) != n {
				n -= 1 << x
				res = append(res, n)
			} else {
				n -= n / 2
				res = append(res, n)
			}
		}

		fmt.Fprintln(out, len(res))
		for i := 0; i < len(res); i++ {
			fmt.Fprint(out, res[i], " ")
		}
		fmt.Fprintln(out)
	}
}
