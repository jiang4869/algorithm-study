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
	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)
	if c >= a {
		fmt.Fprintf(out, "%d-Y ", c)
	} else {
		if d >= b {
			fmt.Fprintf(out, "%d-Y ", c)
		} else {
			fmt.Fprintf(out, "%d-N ", c)
		}
	}
	if d >= a {
		fmt.Fprintf(out, "%d-Y\n", d)
	} else {
		if c >= b {
			fmt.Fprintf(out, "%d-Y\n", d)
		} else {
			fmt.Fprintf(out, "%d-N\n", d)
		}
	}

	if c < a && d >= b {
		fmt.Fprintf(out, "qing %d zhao gu hao %d", 2, 1)
	} else if d < a && c >= b {
		fmt.Fprintf(out, "qing %d zhao gu hao %d", 1, 2)
	} else if c >= a && d >= a {
		fmt.Fprintf(out, "huan ying ru guan")
	} else if c < a && d < a {
		fmt.Fprintln(out, "zhang da zai lai ba")
	} else if c < a && d >= a {
		fmt.Fprintf(out, "%d: huan ying ru guan", 2)
	} else if c >= a && d < a {
		fmt.Fprintf(out, "%d: huan ying ru guan", 1)
	}

}
