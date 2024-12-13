package main

func minimumRounds(tasks []int) int {
	mmid := make(map[int]int)
	for _, task := range tasks {
		if val, ok := mmid[task]; ok {
			mmid[task] = val + 1
		} else {
			mmid[task] = 1
		}
	}

	ans := 0
	for _, val := range mmid {
		if val < 2 {
			return -1
		}
		if val%3 == 0 {
			ans += val / 3
		} else if val%3 == 2 {
			ans += val/3 + 1
		} else if val%3 == 1 {
			ans += (val-3)/3 + 2
		}
	}
	return ans
}
