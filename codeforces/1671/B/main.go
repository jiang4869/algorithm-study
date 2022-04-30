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
	fmt.Fscan(in, &t, "\n")
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		flag := false
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				tmp := make([]int, n)
				copy(tmp, arr)
				tmp[0] += i
				tmp[n-1] += j
				f := true
				for k := 1; k < n-1; k++ {
					want := tmp[k-1] + 1
					if want == tmp[k] {
						continue
					}
					if want == tmp[k]+1 {
						tmp[k] += 1
					} else if want == tmp[k]-1 {
						tmp[k] -= 1
					} else {
						f = false
						break
					}
				}
				if n > 1 && tmp[n-1]-1 != tmp[n-2] {
					f = false
				}
				flag = flag || f
			}
		}

		if flag {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
