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
		var n, d int
		fmt.Fscan(in, &n, &d)
		maxTimes := 120
		if n < 5 {
			tmp := 1
			for i := 1; i <= n; i++ {
				tmp *= i
			}
			maxTimes = tmp
		}
		res := []int{1}
		for i := 3; i < 10; i += 2 {
			num := d
			cur := 1
			flag := false
			for {
				if num%i == 0 {
					flag = true
					break
				} else {
					num = num*10 + d
					cur++
				}
				if cur > maxTimes || cur > 100 {
					break
				}
			}
			if !flag {
				continue
			}
			flag = false
		}
		for _, val := range res {
			fmt.Fprintf(out, "%d ", val)
		}
		fmt.Fprintln(out, "")
	}
}
