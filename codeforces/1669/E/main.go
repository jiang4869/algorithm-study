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
		var n int
		fmt.Fscanf(in, "%d\n", &n)
		mmid := make(map[uint8]map[uint8]int)
		for i := 'a'; i <= 'k'; i++ {
			mmid[uint8(i)] = make(map[uint8]int)
		}
		var str string
		res := int64(0)

		for i := 0; i < n; i++ {
			fmt.Fscanf(in, "%s\n", &str)

			for k := 'a'; k <= 'k'; k++ {
				if str[0] == uint8(k) {
					continue
				}
				res += int64(mmid[uint8(k)][str[1]])
			}
			for k := 'a'; k <= 'k'; k++ {
				if str[1] == uint8(k) {
					continue
				}
				res += int64(mmid[str[0]][uint8(k)])
			}
			mmid[str[0]][str[1]]++

		}

		fmt.Fprintln(out, res)
	}
}
