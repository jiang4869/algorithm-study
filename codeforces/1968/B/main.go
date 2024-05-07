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
	fmt.Fscanf(in, "%d\n", &t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Fscanf(in, "%d%d\n", &n, &m)
		var a, b string
		fmt.Fscanf(in, "%s\n", &a)
		fmt.Fscanf(in, "%s\n", &b)
		ans := 0
		i := 0
		j := 0
		for ; i < len(a); i++ {
			for ; j < len(b); j++ {
				if b[j] == a[i] {
					ans++
					j++
					break
				}
			}
			if j == len(b) {
				break
			}
		}
		fmt.Fprintln(out, ans)
	}
}
