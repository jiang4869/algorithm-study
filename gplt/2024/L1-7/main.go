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
	var a, b int
	fmt.Fscan(in, &a, &b)
	mmid := make(map[int]int)
	for i := a; i <= b; i++ {
		tmpNum := i
		arr := make([]int, 0)
		for {
			tmp := 1
			for tmpNum > 0 {
				tmp *= tmpNum % 10
				tmpNum /= 10
			}
			if count, ok := mmid[tmp]; ok {
				mmid[i] = count + len(arr)
				for j := 0; j < len(arr); j++ {
					mmid[arr[j]] = count + len(arr) - j
				}
				break
			}
			arr = append(arr, tmp)
			if tmp < 10 {
				for j := 0; j < len(arr); j++ {
					mmid[arr[j]] = len(arr) - j
				}
				if i >= 10 {
					mmid[i] = len(arr)
				} else {
					mmid[i] = 0
				}
				break
			}
			tmpNum = tmp
		}
	}
	max := 0
	for _, val := range mmid {
		if max < val {
			max = val
		}
	}
	res := make([]int, 0)
	for k, val := range mmid {
		if val == max && k >= a && k <= b {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	fmt.Fprintln(out, max)
	for i := 0; i < len(res); i++ {
		if i != 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, res[i])
	}
}
