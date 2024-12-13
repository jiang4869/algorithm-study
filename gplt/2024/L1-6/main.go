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
	arr := [120]int{}
	prices := [120]float64{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &prices[i+1])
	}
	var num, count int
	for {
		fmt.Fscan(in, &num, &count)
		if count == 0 {
			break
		}
		arr[num] += count
	}
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += float64(arr[i]) * prices[i]
		fmt.Fprintln(out, arr[i])
	}
	fmt.Fprintf(out, "%.2f", sum)
}
