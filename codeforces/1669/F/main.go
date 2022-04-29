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
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int64, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		total := 0
		leftSum := int64(arr[0])
		rightSum := int64(arr[n-1])
		left, right := 0, n-1

		for left < right {
			if leftSum == rightSum {
				if total < left+(n-right+1) {
					total = left + (n - right + 1)
				}
			}
			if leftSum <= rightSum {
				left++
				leftSum += arr[left]
			} else {
				right--
				rightSum += arr[right]
			}
		}
		fmt.Fprintln(out, total)
	}

}
