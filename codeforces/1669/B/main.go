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
		var num int
		mmid := make(map[int]int)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &num)
			mmid[num]++
		}
		flag := true
		for k, v := range mmid {
			if v >= 3 {
				flag = false
				fmt.Fprintln(out, k)
				break
			}
		}
		if flag {
			fmt.Fprintln(out, -1)
		}

	}
}
