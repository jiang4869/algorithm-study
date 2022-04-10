package main

import (
	"sort"
	"strconv"
)

func largestInteger(num int) int {
	formatInt := strconv.FormatInt(int64(num), 10)
	oddidx := make([]int, 0)
	oddval := make([]int, 0)
	evenidx := make([]int, 0)
	evenval := make([]int, 0)
	for idx, val := range formatInt {
		if (val-'0')%2 == 0 {
			oddidx = append(oddidx, idx)
			oddval = append(oddval, int(val-'0'))
		} else {
			evenidx = append(evenidx, idx)
			evenval = append(evenval, int(val-'0'))
		}
	}
	sort.Slice(oddval, func(i, j int) bool {
		return oddval[i] > oddval[j]
	})
	sort.Slice(evenval, func(i, j int) bool {
		return evenval[i] > evenval[j]
	})
	res := make([]int, len(formatInt))
	for idx, val := range oddidx {
		res[val] = oddval[idx]
	}
	for idx, val := range evenidx {
		res[val] = evenval[idx]
	}

	num = 0
	for _, val := range res {
		num = num*10 + val
	}

	return num
}
