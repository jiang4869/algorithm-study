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
	const mod = 998244353
	var n int
	fmt.Fscan(in, &n)
	var str string
	fmt.Fscan(in, &str)
	str = " " + str

	arr := make([]string, 2<<19)
	res := 1
	var dfs func(rt, length int)
	dfs = func(rt, length int) {
		if length == 1 {
			arr[rt] = string(str[rt])
			return
		}

		dfs(rt<<1, length>>1)
		dfs(rt<<1|1, length>>1)
		a := arr[rt<<1]
		b := arr[rt<<1|1]
		if a != b {
			res = (res + res) % mod
		}
		if a > b {
			arr[rt<<1], arr[rt<<1|1] = arr[rt<<1|1], arr[rt<<1]
		}
		arr[rt] = string(str[rt]) + arr[rt<<1] + arr[rt<<1|1]
	}
	dfs(1, 1<<n-1)
	fmt.Fprintln(out, res%mod)
}
