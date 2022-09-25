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

type Pair struct {
	First  int
	Second int
}

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		if n == 1 {
			fmt.Fprintln(out, "0")
			continue
		}
		res := make([]Pair, 0)
		res = append(res, Pair{
			First:  1,
			Second: n,
		})
		if (arr[0]+arr[n-1])%2 == 0 {
			arr[0] = arr[n-1]

		} else {
			arr[n-1] = arr[0]
		}
		for i := 1; i < n-1; i++ {
			if (arr[0]+arr[i])%2 == 0 {
				res = append(res, Pair{
					First:  i + 1,
					Second: n,
				})
			} else {
				res = append(res, Pair{
					First:  1,
					Second: i + 1,
				})
			}
		}
		fmt.Fprintln(out, len(res))
		for _, val := range res {
			fmt.Fprintln(out, val.First, val.Second)
		}
	}
}

/*

1 10000 3 0 5
1 10000 3 0 1
1 1 3 0 1
1 1 1 0 1
1 1 1 1 1

*/
