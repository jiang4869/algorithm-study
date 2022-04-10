package main

import "sort"

func maximumBeauty(flowers []int, newFlowers int64, target, full, partial int) int64 {
	var res int64
	nums := make([]int, len(flowers))
	copy(nums, flowers)
	sort.Slice(flowers, func(i, j int) bool {
		return flowers[i] > flowers[j]
	})
	pre := make([]int64, len(flowers)+10)
	size := len(flowers)
	if flowers[size-1] >= target {
		res = int64(size) * int64(full)
		return res
	}

	sort.Ints(nums)
	for idx := range nums {
		pre[idx+1] = int64(nums[idx]) + pre[idx]
	}

	judge := func(nf int64, mid int, pos int) bool {
		p := lowerBound(nums, 0, pos, mid)
		var tmp int64
		tmp = int64(p)*int64(mid) - pre[p]
		return tmp <= nf
	}

	for idx, val := range flowers {
		if val >= target {
			continue
		}
		l, r := flowers[size-1], target-1
		tmp := 0
		for l <= r {
			mid := l + (r-l)/2
			if judge(newFlowers, mid, size-idx) {
				tmp = int(max(int64(mid), int64(tmp)))
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		res = max(res, int64(full)*int64(idx)+int64(partial)*int64(tmp))
		newFlowers -= int64(target - val)
		if newFlowers <= 0 {
			return res
		}
	}

	if newFlowers >= 0 {
		res = max(res, int64(full)*int64(size))
	}

	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func lowerBound(nums []int, begin, end int, key int) int {
	l, r := begin, end
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= key {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
