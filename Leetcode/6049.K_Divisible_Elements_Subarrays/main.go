package main

import (
	"fmt"
)

func countDistinct(nums []int, k int, p int) int {
	mmid := make(map[string]int)
	length := len(nums)

	for i := 0; i < length; i++ {
		end := i
		cnt := 0
		res := ""
		for end < length {
			if nums[end]%p == 0 {
				cnt++
			}
			res += fmt.Sprintf("%03d", nums[end])
			if cnt <= k {
				mmid[res]++
			} else {
				break
			}
			end++
		}
	}
	return len(mmid)
}
