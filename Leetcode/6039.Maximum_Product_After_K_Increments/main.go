package main

import (
	"container/heap"
	"sort"
)

type IntHeap struct {
	sort.IntSlice
}

func (h IntHeap) Push(x interface{}) {

}
func (h IntHeap) Pop() interface{} {
	return nil
}

const mod = int(1e9 + 7)

func maximumProduct(nums []int, k int) int {
	que := IntHeap{nums}
	heap.Init(&que)
	for k > 0 {
		que.IntSlice[0]++
		heap.Fix(&que, 0)
		k--
	}
	res := 1
	for _, val := range que.IntSlice {
		res = res * val % mod
	}
	return res
}
