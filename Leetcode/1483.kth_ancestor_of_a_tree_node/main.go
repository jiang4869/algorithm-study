package main

import "math/bits"

type TreeAncestor [][]int

func Constructor(n int, parent []int) TreeAncestor {
	m := bits.Len(uint(n))
	pa := make([][]int, n)
	for i, p := range parent {
		pa[i] = make([]int, m)
		pa[i][0] = p
	}
	for i := 0; i < m-1; i++ {
		for x := 0; x < n; x++ {
			if p := pa[x][i]; p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}
	return pa
}

func (this TreeAncestor) GetKthAncestor(node int, k int) int {
	m := bits.Len(uint(k))
	for i := 0; i < m; i++ {
		if k>>i&1 > 0 {
			node = this[node][i]
			if node == -1 {
				return node
			}
		}
	}
	return node
}
