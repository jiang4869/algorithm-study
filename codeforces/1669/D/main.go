package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	var t int
	fmt.Fscanf(in, "%d\n", &t)
	for ; t > 0; t-- {
		var str string
		var n int
		fmt.Fscanf(in, "%d\n", &n)
		fmt.Fscanf(in, "%s\n", &str)
		splits := strings.Split(str, "W")
		res := true
		for _, val := range splits {
			if len(val) > 0 {
				cnt := 0
				for _, ch := range val {
					if ch == 'B' {
						cnt++
					}
				}
				if cnt == 0 || cnt == len(val) {
					res = false
				}
			}
			if !res {
				break
			}
		}
		if res {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
