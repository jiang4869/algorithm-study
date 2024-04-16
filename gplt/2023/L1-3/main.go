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
	var n, m, k int
	var x string
	fmt.Fscan(in, &n, &x, &m, &k)
	if k == n {
		fmt.Fprintf(out, "mei you mai %s de", x)
	} else if k == m {
		fmt.Fprintf(out, "kan dao le mai %s de", x)
	} else {
		fmt.Fprintf(out, "wang le zhao mai %s de", x)
	}
}
