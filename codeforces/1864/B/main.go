package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var n, k int
		fmt.Fscanf(in, "%d %d\n", &n, &k)
		var s string
		fmt.Fscanf(in, "%s\n", &s)
		str := []byte(s)
		if k%2 == 0 {
			sort.Slice(str, func(i, j int) bool {
				return str[i] < str[j]
			})
			fmt.Fprintln(out, string(str))
		} else {
			a := make([]byte, 0)
			b := make([]byte, 0)
			for i := 0; i < len(str); i++ {
				if i%2 == 0 {
					a = append(a, str[i])
				} else {
					b = append(b, str[i])
				}
			}
			sort.Slice(a, func(i, j int) bool {
				return a[i] < a[j]
			})
			sort.Slice(b, func(i, j int) bool {
				return b[i] < b[j]
			})
			for i := 0; i < len(str); i++ {
				if i%2 == 0 {
					str[i] = a[0]
					a = a[1:]
				} else {
					str[i] = b[0]
					b = b[1:]
				}
			}
			fmt.Fprintln(out, string(str))
		}
	}
}
