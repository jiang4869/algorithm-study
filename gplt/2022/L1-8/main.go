package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

type Student struct {
	GPLT int
	PAT  int
}

func NewStudent(a, b int) Student {
	return Student{GPLT: a, PAT: b}
}

func main() {
	defer out.Flush()
	var n, k, s int
	fmt.Fscan(in, &n, &k, &s)
	arr := make([]Student, n)
	mmid := make(map[int][]Student)

	for i := 0; i < n; i++ {
		var gplt, pat int
		fmt.Fscan(in, &gplt, &pat)
		arr[i] = NewStudent(gplt, pat)
		mmid[gplt] = append(mmid[gplt], arr[i])
	}
	res := 0
	keys := make([]int, 0)
	for key, _ := range mmid {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]

	})

	for _, key := range keys {
		if key >= 175 {
			val := mmid[key]
			tmp := 0
			for _, st := range val {
				if st.PAT >= s {
					tmp++
				}
			}
			res += min(tmp+k, len(val))
		}
	}
	fmt.Fprintln(out, res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
