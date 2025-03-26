package main

type node struct {
	end   int
	begin int
}

const INF = 1e9 + 10

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func longestConsecutive(nums []int) int {
	ans := 0
	mmid := make(map[int]struct{})
	for _, val := range nums {
		mmid[val] = struct{}{}
	}
	for k := range mmid {
		if _, ok := mmid[k-1]; ok {
			continue
		}
		tmp := 1
		num := k
		for {
			if _, ok := mmid[num+1]; ok {
				num = num + 1
				tmp++
			} else {
				break
			}
		}
		ans = max(tmp, ans)
	}
	return ans
}
