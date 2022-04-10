package main

import "sort"

func findWinners(matches [][]int) [][]int {
	winner := make(map[int]int)
	loser := make(map[int]int)

	for _, val := range matches {
		if _, ok := loser[val[0]]; !ok {
			winner[val[0]] = 1
		}

		if _, ok := winner[val[1]]; ok {
			delete(winner, val[1])
		}
		i := loser[val[1]]
		loser[val[1]] = i + 1
	}

	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	for key, _ := range winner {
		arr1 = append(arr1, key)
	}
	for key, val := range loser {
		if val == 1 {
			arr2 = append(arr2, key)
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	return [][]int{
		arr1,
		arr2,
	}
}
