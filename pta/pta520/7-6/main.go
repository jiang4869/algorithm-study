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
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	r := int(float64(n) / 2.718)
	a := 0
	ans1 := 0
	ans2 := 0
	for i := 0; i < r; i++ {
		if arr[i] > a {
			a = arr[i]
		}
	}
	for i := r; i < n; i++ {
		if arr[i] > a {
			ans1 = i + 1
			break
		}
	}
	for i := 1; i < n; i++ {
		if arr[i] > arr[ans2] {
			ans2 = i
		}
	}
	fmt.Fprintln(out, ans1, ans2+1)
}
