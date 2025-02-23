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
		arr := make([]int, n-2)
		for i := 0; i < n-2; i++ {
			fmt.Fscan(in, &arr[i])
		}
		ans := "YES"
		lastIndex := -1
		for i := 0; i < len(arr); i++ {
			if arr[i] == 1 {
				if lastIndex == -1 {
					lastIndex = i
				} else {
					count := i - lastIndex
					lastIndex = i
					if count == 2 {
						ans = "NO"
						break
					}
				}
			}

		}
		fmt.Fprintln(out, ans)
	}
}

/*

10
0 1 0 1 1 0 0 1

1 2 2 2 3 3 3

*/
