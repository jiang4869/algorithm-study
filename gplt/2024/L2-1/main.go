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

// 两个比较大的测试点一直超时，对于go又没有额外的时间，估计用go是不能过了。用c++一样的逻辑就过了。
func main() {
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	mmid := make(map[int]map[int]struct{})
	for i := 1; i <= m; i++ {
		mmid[i] = make(map[int]struct{})
	}
	for i := 1; i <= n; i++ {
		var k int
		fmt.Fscan(in, &k)
		var num int
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &num)
			mmid[num][i] = struct{}{}
		}
	}
	var q int
	fmt.Fscan(in, &q)
	var a, b int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &a, &b)
		m1 := mmid[a]
		m2 := mmid[b]
		if m1 == nil || m2 == nil {
			fmt.Fprintln(out, "0")
			continue
		}
		count := 0
		for k := range m1 {
			if _, ok := m2[k]; ok {
				count++
			}
		}
		fmt.Fprintln(out, count)
	}
}
