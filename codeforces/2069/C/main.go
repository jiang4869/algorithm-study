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

const mod = 998244353

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		var num int64
		a, b, c := 0, 0, 0
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &num)
			if num == 1 {
				a++
			} else if num == 2 {
				b = (b*2 + a) % mod
			} else {
				c += b
				c %= mod
			}

		}
		fmt.Fprintln(out, c%mod)
	}
}
