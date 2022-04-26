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

	var a, b int
	res := 1
	fmt.Fscan(in, &a, &b)

	for i := 1; i <= a+b; i++ {
		res = res * i
	}

	fmt.Fprintln(out, res)
}
