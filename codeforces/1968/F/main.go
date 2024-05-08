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
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, q int
		fmt.Fscan(in, &n, &q)
		mmid := make(map[int][]int)
		mmid[0] = []int{0}
		arr := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &arr[i])
			arr[i] ^= arr[i-1]
			if val, ok := mmid[arr[i]]; ok {
				val = append(val, i)
				mmid[arr[i]] = val
			} else {
				mmid[arr[i]] = []int{i}
			}
		}
		for i := 0; i < q; i++ {
			var l, r int
			fmt.Fscan(in, &l, &r)
			if arr[l-1] == arr[r] {
				fmt.Fprintln(out, "YES")
				continue
			}
			a1 := mmid[arr[r]]
			it1 := sort.Search(len(a1), func(i int) bool {
				return a1[i] >= l
			})
			if it1 == len(a1) {
				fmt.Fprintln(out, "NO")
				continue
			}
			t1 := a1[it1]
			a2 := mmid[arr[l-1]]
			it2 := sort.Search(len(a2), func(i int) bool {
				return a2[i] > t1
			})
			if it2 == len(a2) {
				fmt.Fprintln(out, "NO")
				continue
			}
			t2 := a2[it2]
			if t1 < r && t2 < r {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
		}
		fmt.Fprintln(out, "")
	}
}
