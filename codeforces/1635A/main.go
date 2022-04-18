package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  *bufio.Reader = bufio.NewReader(os.Stdin)
	out *bufio.Writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		var n int
		fmt.Fscan(in, &n)
		ans := 0
		for i := 0; i < n; i++ {
			var x int
			fmt.Fscan(in, &x)
			ans |= x
		}
		fmt.Fprintln(out, ans)
		t--
	}
}
