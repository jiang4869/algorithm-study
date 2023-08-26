package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化后的版本
func longestEqualSubarray(nums []int, k int) int {
	type pair struct{ idx, val int }
	v := make([][]int, len(nums)+1)
	poss := make([]int, len(nums)+1)
	for idx, val := range nums {
		if poss[val] == 0 {
			poss[val] = idx + 1
		} else {
			pos := poss[val] - 1
			v[val] = append(v[val], idx-pos-1)
			poss[val] = idx + 1
		}
	}
	ans := 1
	for _, arr := range v {
		begin := 0
		end := 0
		kk := k
		sum := 0
		for ; end < len(arr); end++ {
			sum += arr[end]
			for sum > kk {
				sum -= arr[begin]
				begin++
			}
			ans = max(ans, end-begin+2)
		}
	}
	return ans
}

//func longestEqualSubarray(nums []int, k int) int {
//	type pair struct{ idx, val int }
//	v := make([][]pair, len(nums)+1)
//	for idx, val := range nums {
//		v[val] = append(v[val], pair{idx: idx, val: val})
//	}
//	ans := 1
//	for _, arr := range v {
//		begin := 0
//		end := 0
//		kk := k
//		tmp := make([]int, len(arr))
//		for end < len(arr) && kk >= 0 {
//			end++
//			if end < len(arr) && kk-(arr[end].idx-arr[end-1].idx-1) >= 0 {
//				kk = kk - (arr[end].idx - arr[end-1].idx - 1)
//				tmp[end] = arr[end].idx - arr[end-1].idx - 1
//				ans = max(ans, end-begin+1)
//			} else {
//				for begin < end {
//					begin++
//					if begin < len(arr) {
//						kk = kk + tmp[begin]
//						if end < len(arr) && kk-(arr[end].idx-arr[end-1].idx-1) >= 0 {
//							kk = kk - (arr[end].idx - arr[end-1].idx - 1)
//							tmp[end] = arr[end].idx - arr[end-1].idx - 1
//							ans = max(ans, end-begin+1)
//							break
//						}
//					}
//				}
//			}
//		}
//	}
//	return ans
//}
