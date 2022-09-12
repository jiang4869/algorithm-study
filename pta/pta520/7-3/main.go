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
	var sex, a, b, c, d int
	fmt.Fscan(in, &sex, &a, &b, &c, &d)
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var s, age, height int
		fmt.Fscan(in, &s, &age, &height)
		if sex != s && age >= a && age <= b && height >= c && height <= d {
			fmt.Fprintln(out, s, age, height)
		}
	}
}
