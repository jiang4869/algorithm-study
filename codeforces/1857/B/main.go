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
	var (
		t   int
		str string
	)
	fmt.Fscanf(in, "%d\n", &t)
	for ; t > 0; t-- {
		fmt.Fscanf(in, "%s\n", &str)
		str = "0" + str
		n := len(str)
		arr := make([]uint8, len(str))
		for i := 0; i < n; i++ {
			arr[i] = str[i]
		}
		p := n
		for i := n - 1; i >= 0; i-- {
			if arr[i] >= '5' {
				p = i
				arr[i-1]++
			}
		}
		begin := 0
		if arr[0] == '0' {
			begin = 1
		}
		for i := begin; i < n; i++ {
			if i < p {
				fmt.Fprint(out, string(arr[i]))
			} else {
				fmt.Fprint(out, 0)
			}
		}
		fmt.Fprintln(out)
	}
}
