package main

import (
	"bufio"
	"fmt"
	"os"
)

var in *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()
	col := make(map[int]int)
	row := make(map[int]int)
	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	var t, c int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &t, &c)
		if t == 0 {
			row[c] = 1
		} else {
			col[c] = 1
		}
	}
	cnt := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			_, ok1 := row[i]
			_, ok2 := col[j]
			if !ok1 && !ok2 {
				cnt++
			}
		}
	}
	fmt.Fprintln(out, cnt)

}
