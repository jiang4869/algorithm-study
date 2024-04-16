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
	var a, b int
	fmt.Fscan(in, &a, &b)
	sum := a + b
	fmt.Fprintln(out, sum-16)
	fmt.Fprintln(out, sum-3)
	fmt.Fprintln(out, sum-1)
	fmt.Fprintln(out, sum)
}
