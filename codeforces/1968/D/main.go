package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in      = bufio.NewReader(os.Stdin)
	out     = bufio.NewWriter(os.Stdout)
	i64zero = int64(0)
)

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, k, pb, ps int64
		fmt.Fscan(in, &n, &k, &pb, &ps)
		parr := make([]int64, n)
		aarr := make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(in, &parr[i])
			parr[i]--
		}
		for i := int64(0); i < n; i++ {
			fmt.Fscan(in, &aarr[i])
		}
		flag := true
		for i := int64(0); i < n-1; i++ {
			if aarr[i] != aarr[i+1] {
				flag = false
			}
		}
		if flag {
			fmt.Fprintln(out, "Draw")
			continue
		}
		pb--
		ps--

		cal := func(ps int64) int64 {
			sum := i64zero
			ans := i64zero
			for i := i64zero; i < k && i < n; i++ {
				if ans < (k-i)*aarr[ps]+sum {
					ans = (k-i)*aarr[ps] + sum
				}
				sum += aarr[ps]
				ps = parr[ps]
			}
			return ans
		}
		scoreb := cal(pb)
		scores := cal(ps)

		if scoreb > scores {
			fmt.Fprintln(out, "Bodya")
		} else if scoreb < scores {
			fmt.Fprintln(out, "Sasha")
		} else {
			fmt.Fprintln(out, "Draw")
		}
	}
}
