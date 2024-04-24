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
	tip := "-"
	action := "stop"
	if a == 1 {
		action = "move"
	}
	if a == 1 && b == 0 {
		tip = "dudu"
	}
	if a == 0 && b == 0 {
		tip = "biii"
	}
	fmt.Fprintln(out, tip)
	fmt.Fprintln(out, action)
}
