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

type node struct {
	a, b int
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		var num int
		arr := [1010]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &num)
			arr[num]++
		}
		ans := "YES"
		for i := 1; i <= 1000; i++ {
			if arr[i] == 1 {
				ans = "NO"
				break
			}
			if arr[i] > 2 {
				arr[i+1] += arr[i] - 2
			}
		}
		if arr[1000]%2 != 0 {
			ans = "N0"
		}
		fmt.Fprintln(out, ans)
	}
}
