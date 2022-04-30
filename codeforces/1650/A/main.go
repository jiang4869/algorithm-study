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
	fmt.Fscan(in, &t, "\n")
	for ; t > 0; t-- {
		var str string
		var ch string

		fmt.Fscan(in, &str, "\n")
		fmt.Fscan(in, &ch, "\n")
		flag := false
		arr := make([]int, 0)
		for idx, val := range str {
			if uint8(val) == ch[0] {
				arr = append(arr, idx)
			}
		}
		length := len(str)
		for _, val := range arr {
			front := val
			end := length - val - 1
			if (front%2 == 0) && (end%2 == 0) {
				flag = true
			}
		}
		if flag {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
