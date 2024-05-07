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

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var x int
		fmt.Fscan(in, &x)
		ans := 0
		y := 0
		for i := 1; i < x; i++ {
			tmp := gcd(x, i)
			if ans < tmp+i {
				ans = tmp + i
				y = i
			}
		}
		fmt.Fprintln(out, y)
	}
}
