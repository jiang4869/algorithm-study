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

type Node struct {
	Num   int
	Root  bool
	Child []int
}

func main() {
	defer out.Flush()
	var n, root int
	fmt.Fscan(in, &n)
	arr := make([]Node, n+1)
	for i := 0; i < n; i++ {
		arr[i+1].Num = i + 1
	}
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(in, &num)
		if num != 0 {
			arr[num].Child = append(arr[num].Child, i+1)
		} else {
			root = i + 1
		}
	}
	degree := 0
	isK := "yes"
	mmid := make(map[int]int)
	for i := 1; i <= n; i++ {
		degree = max(degree, len(arr[i].Child))
		sort.Ints(arr[i].Child)
		mmid[len(arr[i].Child)] = 0
	}
	if len(mmid) > 2 {
		isK = "no"
	}
	fmt.Fprintf(out, "%d %s\n", degree, isK)
	isFirst := true
	var printTree func(root int)
	printTree = func(root int) {
		if !isFirst {
			fmt.Fprintf(out, " ")
		}
		isFirst = false
		fmt.Fprintf(out, "%d", arr[root].Num)
		for i := 0; i < len(arr[root].Child); i++ {
			printTree(arr[root].Child[i])
		}
	}
	printTree(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
