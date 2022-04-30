package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var n, x int64
		fmt.Fscan(in, &n, &x)
		arr := make([]int64, n)
		for i := int64(0); i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		res := int64(0)
		sum := int64(0)
		for i := int64(0); i < n; i++ {
			sum += arr[i]
			if x >= sum {
				res += (x-sum)/(i+1) + 1
			}
		}

		fmt.Fprintln(out, res)
	}

}
