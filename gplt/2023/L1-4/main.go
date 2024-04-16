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
	for ; n > 0; n-- {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if a+b == c {
			fmt.Fprintln(out, "Tu Dou")
		} else if a*b == c {
			fmt.Fprintln(out, "Lv Yan")
		} else {
			fmt.Fprintln(out, "zhe du shi sha ya!")
		}
	}
}
