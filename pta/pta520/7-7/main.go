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
	var n, m int
	fmt.Fscan(in, &n, &m)

	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &arr[i][j])
		}
	}
	var r, c, num int
	fmt.Fscan(in, &r, &c, &num)
	det := math.MaxInt32
	for i := 1; i <= m; i++ {
		if arr[r][i] != -1 && abs(arr[r][i]-num) < det {
			det = abs(arr[r][i] - num)
		}
	}
	for i := 1; i <= n; i++ {
		if arr[i][c] != -1 && abs(arr[i][c]-num) < det {
			det = abs(arr[i][c] - num)
		}
	}
	for i := 1; i <= m; i++ {
		if arr[r][i] != -1 && abs(arr[r][i]-num) == det {
			fmt.Fprintf(out, "(%d:%d)\n", r, i)
		}
	}
	for i := 1; i <= n; i++ {
		if arr[i][c] != -1 && abs(arr[i][c]-num) == det && i != r {
			fmt.Fprintf(out, "(%d:%d)\n", i, c)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
