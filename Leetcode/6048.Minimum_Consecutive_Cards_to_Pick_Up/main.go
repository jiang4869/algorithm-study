package main

import "math"

func minimumCardPickup(cards []int) int {
	mmid := make(map[int][]int)
	res := math.MaxInt32
	for idx, val := range cards {
		mmid[val] = append(mmid[val], idx)
	}
	for _, v := range mmid {
		length := len(v)
		for i := 0; i < length-1; i++ {
			if (v[i+1] - v[i]) < res {
				res = v[i+1] - v[i] + 1
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
