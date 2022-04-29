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
	fmt.Fscanf(in, "%d\n", &t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Fscanf(in, "%d%d\n", &n, &m)
		var s string
		arr := make([][]uint8, n)
		for i := 0; i < n; i++ {
			arr[i] = make([]uint8, m)
		}
		for i := 0; i < n; i++ {
			fmt.Fscanf(in, "%s\n", &s)
			for j := 0; j < m; j++ {
				arr[i][j] = s[j]
			}
		}

		for k := 0; k < n-1; k++ {

			for i := n - 2; i >= 0; i-- {

				for j := 0; j < m; j++ {
					if arr[i][j] == '.' {
						continue
					}
					if arr[i][j] == 'o' {
						continue
					}
					if arr[i+1][j] == '.' {
						arr[i][j] = '.'
						arr[i+1][j] = '*'
					}
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fprintf(out, "%c", arr[i][j])
			}
			fmt.Fprintln(out)
		}

	}
}
