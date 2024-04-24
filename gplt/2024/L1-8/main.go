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
	for t := 0; t < n; t++ {
		arr := [9][9]int{}
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				fmt.Fscan(in, &arr[j][k])
			}
		}
		flag := true

		for i := 0; i < 9; i++ {
			tmp := [10]int{}
			for j := 0; j < 9; j++ {
				if arr[i][j] >= 1 && arr[i][j] <= 9 {
					tmp[arr[i][j]]++
				}
			}
			for j := 1; j <= 9; j++ {
				if tmp[j] == 0 || tmp[j] > 1 {
					flag = false
					goto out
				}
			}
		}

		for i := 0; i < 9; i++ {
			tmp := [10]int{}
			for j := 0; j < 9; j++ {
				if arr[j][i] >= 1 && arr[j][i] <= 9 {
					tmp[arr[j][i]]++
				}
			}
			for j := 1; j <= 9; j++ {
				if tmp[j] == 0 || tmp[j] > 1 {
					flag = false
					goto out
				}
			}
		}

		for i := 0; i < 9; i += 3 {
			for j := 0; j < 9; j += 3 {
				tmp := [10]int{}

				for k := i; k < i+3; k++ {
					for h := j; h < j+3; h++ {
						if arr[k][h] >= 1 && arr[k][h] <= 9 {
							tmp[arr[k][h]]++
						}
					}
				}

				for j := 1; j <= 9; j++ {
					if tmp[j] == 0 || tmp[j] > 1 {
						flag = false
						goto out
					}
				}
			}
		}
	out:
		if flag {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}
