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
	var n, v int
	fmt.Fscan(in, &n, &v)
	res := n / v
	fmt.Fprintln(out, res)
}
