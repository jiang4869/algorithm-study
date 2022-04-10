package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	len1 := len(nums1)
	len2 := len(nums2)

	totalLeft := int((len1 + len2 + 1) / 2)
	l, r := 0, len1

	for l < r {
		i := l + (r-l+1)/2
		j := totalLeft - i

		if nums1[i-1] > nums2[j] {
			r = i - 1
		} else {
			l = i
		}
	}
	i := l
	j := totalLeft - i
	nums1LeftMax, nums2LeftMax := math.MinInt32, math.MinInt32
	nums1RightMin, nums2RightMin := math.MaxInt32, math.MaxInt32
	if i != 0 {
		nums1LeftMax = nums1[i-1]
	}
	if i != len1 {
		nums1RightMin = nums1[i]
	}
	if j != 0 {
		nums2LeftMax = nums2[j-1]
	}
	if j != len2 {
		nums2RightMin = nums2[j]
	}

	if ((len1 + len2) % 2) == 1 {
		return float64(judge(nums1LeftMax > nums2LeftMax, nums1LeftMax, nums2LeftMax))
	} else {
		return float64(judge(nums1LeftMax > nums2LeftMax, nums1LeftMax, nums2LeftMax)+judge(nums1RightMin > nums2RightMin, nums2RightMin, nums1RightMin)) / 2
	}

}
func judge(flag bool, a, b int) int {
	if flag {
		return a
	} else {
		return b
	}
}
