package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ans)
}
