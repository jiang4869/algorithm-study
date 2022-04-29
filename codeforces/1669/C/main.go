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
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		flag1, flag2 := true, true
		for i := 0; i+2 < n; i++ {
			if (arr[i] % 2) != (arr[i+2] % 2) {
				flag1 = false
				break
			}
		}
		for i := 1; i+2 < n; i++ {
			if (arr[i] % 2) != (arr[i+2] % 2) {
				flag2 = false
				break
			}
		}
		if flag1 && flag2 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
