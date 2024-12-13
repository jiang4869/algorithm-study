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
		var n int
		fmt.Fscan(in, &n)
		arr := [110]int{}
		var num int
		ans := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &num)
			arr[num]++
		}
		for i := 1; i <= 100; i++ {
			ans += arr[i] / 3
		}
		fmt.Fprintln(out, ans)
	}
}
