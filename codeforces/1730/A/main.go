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
		var n, c, num int
		fmt.Fscan(in, &n, &c)
		mmid := make(map[int]int)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &num)
			if val, ok := mmid[num]; ok {
				mmid[num] = val + 1
			} else {
				mmid[num] = 1
			}
		}
		res := 0
		for _, val := range mmid {
			if val >= c {
				res += c
			} else {
				res += val
			}
		}
		fmt.Fprintln(out, res)
	}
}
