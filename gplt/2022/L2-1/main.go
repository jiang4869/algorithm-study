package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var in *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	cur := math.MaxInt32
	st := make([]int, 0)
	num := 0
	mmid := make(map[int][]int)
	for i := 0; i < n; i++ {
		if len(mmid[num]) == k {
			num++
			cur = math.MaxInt32
		}
		if len(st) != 0 {
			if st[len(st)-1] <= cur {
				mmid[num] = append(mmid[num], st[len(st)-1])
				cur = st[len(st)-1]
				st = st[:len(st)-1]
				i--
				continue
			}
		}
		if arr[i] <= cur {
			cur = arr[i]
			mmid[num] = append(mmid[num], arr[i])
		} else {
			if len(st) == m {
				cur = math.MaxInt32
				num++
				i--
			} else {
				st = append(st, arr[i])
			}
		}
	}

	for len(st) > 0 {
		if len(mmid[num]) == k {
			num++
			cur = math.MaxInt32
		}
		if st[len(st)-1] <= cur {
			mmid[num] = append(mmid[num], st[len(st)-1])
			cur = st[len(st)-1]
			st = st[:len(st)-1]
		} else {
			cur = math.MaxInt32
			num++
		}
	}
	num++
	for j := 0; j < num; j++ {
		if j != 0 {
			fmt.Fprintf(out, "\n")
		}
		v := mmid[j]
		for i := 0; i < len(v); i++ {
			if i != 0 {
				fmt.Fprintf(out, " ")
			}
			fmt.Fprintf(out, "%d", v[i])
		}
	}

}
