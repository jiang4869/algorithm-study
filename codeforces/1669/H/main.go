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
		var n, k int
		fmt.Fscan(in, &n, &k)
		arr := make([]int, n)
		bits := make([]int, 32)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
			idx := 0
			for arr[i] > 0 {
				if (arr[i] % 2) == 1 {
					bits[idx]++
				}
				arr[i] /= 2
				idx++
			}
		}
		for i := 30; i >= 0; i-- {
			if bits[i]+k >= n {
				det := n - bits[i]
				if det <= 0 {
					det = 0
				}
				bits[i] = n
				k -= det
			}
		}
		res := 0
		cur := 1
		for i := 0; i < 32; i++ {
			if bits[i] == n {
				res += cur
			}
			cur *= 2
		}
		fmt.Fprintln(out, res)
	}
}
