package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	mmid := make(map[string][]int)
	for i, str := range strs {
		data := []byte(str)
		slices.Sort(data)
		sorted := string(data)
		if val, ok := mmid[sorted]; ok {
			val = append(val, i)
			mmid[sorted] = val
		} else {
			mmid[sorted] = []int{i}
		}
	}
	for _, val := range mmid {
		arr := make([]string, 0)
		for _, i := range val {
			arr = append(arr, strs[i])
		}
		ans = append(ans, arr)
	}

	return ans
}
