package main

import (
	"bufio"
	"fmt"
	"math"
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
		var n, x int
		fmt.Fscan(in, &n, &x)
		arr := make([]int, n)
		minn, maxn := math.MaxInt32, math.MinInt32
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
			minn = min(arr[i], minn)
			maxn = max(arr[i], maxn)
		}
		cur := int64(0)

		if n == 1 {
			if arr[0] >= 1 && arr[0] <= x {
				cur = int64(x - 1)
			} else {
				cur = int64(arr[0] - 1)
			}
		} else {
			for i := 0; i < n-1; i++ {
				cur += int64(abs(arr[i] - arr[i+1]))
			}
			if minn <= 1 && x <= maxn {

			}
			if 1 < minn {
				num1 := abs(min(arr[0], arr[n-1]) - 1)
				num2 := math.MaxInt32
				for i := 0; i < n-1; i++ {
					num2 = min(num2, (min(arr[i], arr[i+1])-1)*2)
				}
				cur += int64(min(num1, num2))
			}
			if x > maxn {
				num1 := abs(x - max(arr[0], arr[n-1]))
				num2 := math.MaxInt32
				for i := 0; i < n-1; i++ {
					num2 = min(num2, (x-max(arr[i], arr[i+1]))*2)
				}
				cur += int64(min(num1, num2))
			}
		}

		fmt.Fprintln(out, cur)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
