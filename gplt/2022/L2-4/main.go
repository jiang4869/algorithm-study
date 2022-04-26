package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

var in *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

type Node struct {
	Dis int
	To  int
}

func NewNode(a, b int) Node {
	return Node{
		Dis: a,
		To:  b,
	}
}

type Pair struct {
	First  int
	Second int
}

func NewPair(a, b int) *Pair {
	return &Pair{
		First:  a,
		Second: b,
	}
}

type PriorityQueue []*Pair

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].First < p[j].First
}

func (p PriorityQueue) Len() int {
	return len(p)
}
func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	item := x.(*Pair)
	*p = append(*p, item)
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // 避免内存泄漏
	*p = old[0 : n-1]
	return item
}

func main() {
	defer out.Flush()
	const INF int = 9e8
	var n int
	fmt.Fscanf(in, "%d\n", &n)
	G := make([][]Node, n+10)
	D := make([][]int, n+10)
	sex := make([]uint8, n+10)
	for i := 0; i <= n; i++ {
		D[i] = make([]int, n+10)
	}
	for i := 1; i <= n; i++ {
		var ch uint8
		var k int
		fmt.Fscanf(in, "%c %d", &ch, &k)
		sex[i] = ch
		for j := 0; j < k; j++ {
			var num, dis int
			fmt.Fscanf(in, "%d:%d", &num, &dis)
			G[i] = append(G[i], NewNode(dis, num))
		}
		fmt.Fscanf(in, "\n")
	}
	dijkstra := func(s int) {
		dis := make([]int, n+10)
		sign := make([]bool, n+10)
		for i := 0; i < n+10; i++ {
			dis[i] = INF
			sign[i] = false
		}
		que := make(PriorityQueue, 0)
		heap.Init(&que)
		heap.Push(&que, NewPair(0, s))
		dis[s] = 0
		for que.Len() > 0 {
			tmp := heap.Pop(&que).(*Pair)
			if sign[tmp.Second] {
				continue
			}
			sign[tmp.Second] = true
			for _, val := range G[tmp.Second] {
				if dis[val.To] > dis[tmp.Second]+val.Dis {
					dis[val.To] = dis[tmp.Second] + val.Dis
					heap.Push(&que, NewPair(dis[val.To], val.To))
				}
			}
		}
		dis[s] = INF
		for i := 1; i <= n; i++ {
			D[i][s] = dis[i]
		}
	}

	for i := 1; i <= n; i++ {
		dijkstra(i)
	}

	maxs := make([]int, n+10)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if sex[i] != sex[j] {
				maxs[i] = max(maxs[i], D[i][j])
			}
		}
	}
	max1, max2 := math.MaxInt32, math.MaxInt32

	for i := 1; i <= n; i++ {
		if sex[i] == 'F' {
			max1 = min(max1, maxs[i])
		} else {
			max2 = min(max2, maxs[i])
		}
	}
	flag := false
	for i := 1; i <= n; i++ {
		if sex[i] == 'F' && maxs[i] == max1 {
			if flag {
				fmt.Fprintf(out, " ")
			}
			flag = true
			fmt.Fprintf(out, "%d", i)
		}
	}
	fmt.Fprintln(out)
	flag = false
	for i := 1; i <= n; i++ {
		if sex[i] == 'M' && maxs[i] == max2 {
			if flag {
				fmt.Fprintf(out, " ")
			}
			flag = true
			fmt.Fprintf(out, "%d", i)
		}
	}
	fmt.Fprintln(out)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
