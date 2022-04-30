package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  *bufio.Reader = bufio.NewReader(os.Stdin)
	out *bufio.Writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t, "\n")
	for ; t > 0; t-- {
		var str string
		fmt.Fscan(in, &str, "\n")

		i := 0
		slen := len(str)
		flag := true
		for ; i < slen; i++ {
			cnt := 0
			if str[i] == 'a' {
				for ; i < slen; i++ {
					if str[i] != 'a' {
						break
					}
					cnt++
				}
				i--
				if cnt < 2 {
					flag = false
				}
			} else {
				for ; i < slen; i++ {
					if str[i] != 'b' {
						break
					}
					cnt++
				}
				i--
				if cnt < 2 {
					flag = false
				}
			}
		}
		if flag {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
