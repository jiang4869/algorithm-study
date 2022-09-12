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
	fmt.Fscan(in, &n, &m, &k)
	if n*k == m {
		fmt.Fprintf(out, "zheng hao mei ren %d!", k)
	} else if n*k < m {
		fmt.Fprintf(out, "hai sheng %d!", m-n*k)
	} else {
		fmt.Fprintf(out, "hai cha %d!", n*k-m)
	}
}
