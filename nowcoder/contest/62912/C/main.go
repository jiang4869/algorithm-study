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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://ac.nowcoder.com/acm/contest/62912/C
// 对公式进行拆开，提取换算一下，可以得到该式子等价于所有数之和*min(i,n-1)之和
func main() {
	defer out.Flush()
	var (
		n   int
		num int64
	)
	mod := int64(1e9 + 7)
	a := int64(0)
	b := a
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &num)
		a = (a + num) % mod
		b = (b + int64(min(i, n-i))) % mod
	}
	fmt.Fprintln(out, (a*b)%mod)
}
