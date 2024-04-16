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
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	var k int
	fmt.Fscan(in, &k)
	for i := 0; i < k; i++ {
		f := 0
		t := 0
		num := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &num)
			if num == 0 {
				continue
			}
			if arr[j] == num {
				t++
			} else {
				f++
			}
		}
		if t > 0 && f == 0 {
			fmt.Fprintln(out, "Da Jiang!!!")
		} else {
			fmt.Fprintln(out, "Ai Ya")
		}
	}
}
