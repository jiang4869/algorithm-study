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
	var n, m int
	fmt.Fscan(in, &n, &m)
	flag := false
	for x := 1; x <= m; x += 2 {
		for y := x + 2; y <= m; y += 2 {
			for z := y + 2; z <= m; z += 2 {
				if 3*x*y*z == n*(z*y+x*z+y*x) {
					flag = true
					fmt.Fprintln(out, x, y, z)
				}
				if flag {
					break
				}
			}
			if flag {
				break
			}
		}
		if flag {
			break
		}
	}
	if !flag {
		fmt.Fprintf(out, "No solution in (3, %d].", m)
	}
}
