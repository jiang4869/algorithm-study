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
	var n, m int
	fmt.Fscan(in, &n, &m)
	min := int(1e9)
	var sum, num int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &num)
		sum += num
		if min > num {
			min = num
		}
	}
	ans := sum - n*(m-1)
	if ans < 0 {
		ans = 0
	}
	fmt.Fprintln(out, ans)
}
