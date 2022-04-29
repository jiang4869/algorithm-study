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
		var rating int
		fmt.Fscan(in, &rating)
		div := 1
		if rating >= 1900 {
			div = 1
		} else if 1600 <= rating && rating <= 1899 {
			div = 2
		} else if 1400 <= rating && rating <= 1599 {
			div = 3
		} else {
			div = 4
		}
		fmt.Fprintln(out, "Division", div)
	}

}
