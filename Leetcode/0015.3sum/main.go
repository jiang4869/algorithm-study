package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		thirdIndex := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j]
			for thirdIndex > j && nums[thirdIndex]+sum > 0 {
				thirdIndex--
			}
			if thirdIndex == j {
				break
			}
			if sum+nums[thirdIndex] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[thirdIndex]})
			}
		}
	}

	return ans
}

/*



func toString(a, b, c int) string {
	return fmt.Sprintf("%d,%d,%d", a, b, c)
}

func threeSum(nums []int) [][]int {
	mmid := make(map[int][]int)
	sort.Ints(nums)
	existed := make(map[string]struct{})
	for i, val := range nums {
		if arr, ok := mmid[val]; ok {
			arr = append(arr, i)
			mmid[val] = arr
		} else {
			mmid[val] = []int{i}
		}
	}
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if arr, ok := mmid[-sum]; ok {
				val := arr[len(arr)-1]
				if val > j {
					a := []int{nums[i], nums[j], nums[val]}
					s := toString(a[0], a[1], a[2])
					if _, find := existed[s]; !find {
						ans = append(ans, a)
						existed[s] = struct{}{}
					}
				}

			}
		}
	}

	return ans
}

*/
