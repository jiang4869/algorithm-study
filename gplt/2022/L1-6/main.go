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
	var str1, str2 string
	fmt.Fscan(in, &str1, &str2)
	var res1, res2 string
	for i := 1; i < len(str1); i++ {
		if str1[i]%2 == str1[i-1]%2 {
			res1 += string(max(str1[i], str1[i-1]))
		}
	}
	for i := 1; i < len(str2); i++ {
		if str2[i]%2 == str2[i-1]%2 {
			res2 += string(max(str2[i], str2[i-1]))
		}
	}

	if res1 == res2 {
		fmt.Fprintln(out, res1)
	} else {
		fmt.Fprintln(out, res1)
		fmt.Fprintln(out, res2)
	}
}

func max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
