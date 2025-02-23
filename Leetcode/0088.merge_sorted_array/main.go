package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	a, b, c := n+m-1, m-1, n-1
	for a >= 0 {
		if b >= 0 && c >= 0 {
			if nums1[b] > nums2[c] {
				nums1[a] = nums1[b]
				b--
			} else {
				nums1[a] = nums2[b]
				c--
			}
		} else if b >= 0 {
			nums1[a] = nums1[b]
			b--
		} else {
			nums1[a] = nums2[b]
			c--
		}
		a--
	}
}
