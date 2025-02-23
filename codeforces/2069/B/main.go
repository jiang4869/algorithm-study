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
		var n, m int
		fmt.Fscan(in, &n, &m)
		arr := make([][]int, n)
		mmid := make(map[int]int)
		for i := 0; i < n; i++ {
			arr[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &arr[i][j])
				if _, ok := mmid[arr[i][j]]; !ok {
					mmid[arr[i][j]] = 1
				}
				if i > 0 && arr[i][j] == arr[i-1][j] {
					mmid[arr[i][j]] = 2
				}
				if j > 0 && arr[i][j] == arr[i][j-1] {
					mmid[arr[i][j]] = 2
				}
			}
		}
		count1, count2 := 0, 0
		for _, val := range mmid {
			count1 += val
			if count2 < val {
				count2 = val
			}
		}
		fmt.Fprintln(out, count1-count2)
	}
}
