package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()
	arr := make([]int, 6)
	var n int
	mmid := make(map[int]map[int]int)

	for i := 0; i < 6; i++ {
		fmt.Fscan(in, &arr[i])
		mmid[i] = make(map[int]int)
		mmid[i][arr[i]] = 1
	}
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		for k := 0; k < 6; k++ {
			for j := 6; j >= 1; j-- {
				if _, ok := mmid[k][j]; !ok {
					arr[k] = j
					mmid[k][j] = 1
					break
				}
			}
		}
	}
	for i := 0; i < 6; i++ {
		if i != 0 {
			fmt.Fprintf(out, " ")
		}
		fmt.Fprintf(out, "%d", arr[i])
	}
}
