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

func count(a, b string) []int {
	arr := make([]int, 10)
	begin := 0
	for ; begin < len(a); begin++ {
		if a[begin] != '0' {
			break
		}
	}
	for ; begin < len(a); begin++ {
		arr[a[begin]-'0']++
	}
	begin = 0
	for ; begin < len(b); begin++ {
		if a[begin] != '0' {
			break
		}
	}
	for ; begin < len(b); begin++ {
		arr[b[begin]-'0']++
	}
	return arr
}

func main() {
	defer out.Flush()
	var a, b string
	fmt.Fscan(in, &a, &b, "\n")
	arr := count(a, b)
	for {
		var c, d string
		fmt.Fscan(in, &c, &d, "\n")
		if c == "0" && d == "0" {
			break
		}
		tmp := count(c, d)
		flag := true
		for i := 0; i < 10; i++ {
			if arr[i] != tmp[i] {
				flag = false
				break
			}
		}
		if flag {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
