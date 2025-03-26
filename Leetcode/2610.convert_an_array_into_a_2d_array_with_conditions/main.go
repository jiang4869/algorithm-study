package main

func findMatrix(nums []int) [][]int {
	ans := make([][]int, 0)
	mmid := make(map[int]int)
	for _, num := range nums {
		mmid[num]++
	}
	for len(mmid) > 0 {
		arr := make([]int, 0)
		for k := range mmid {
			arr = append(arr, k)
			mmid[k]--
		}
		ans = append(ans, arr)
		for k, v := range mmid {
			if v == 0 {
				delete(mmid, k)
			}
		}
	}
	return ans
}
