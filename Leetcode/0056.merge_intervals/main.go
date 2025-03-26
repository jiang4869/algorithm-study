package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	var interval []int
	for _, arr := range intervals {
		if interval == nil {
			interval = make([]int, 0)
			interval = append(interval, arr...)
			continue
		}
		if interval[1] >= arr[0] {
			interval[1] = arr[1]
		} else {
			ans = append(ans, interval)
			interval = make([]int, 0)
			interval = append(interval, arr...)
		}
	}
	ans = append(ans, interval)
	return ans
}
