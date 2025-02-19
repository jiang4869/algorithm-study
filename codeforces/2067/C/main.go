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
	} else {
		return b
	}
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		ans := 7
		for i := 0; i < 10; i++ {

			str := fmt.Sprintf("%d", n-i)
			for _, val := range str {
				num := int(val - '0')
				if num <= 7 && i+num >= 7 {
					ans = min(ans, i)
				}
			}

		}
		fmt.Fprintln(out, ans)
	}
}
