package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	subarray := findUnsortedSubarray(nums)
	fmt.Println(subarray)
}

func Test_Case2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	subarray := findUnsortedSubarray(nums)
	fmt.Println(subarray)
}

func Test_Case3(t *testing.T) {
	nums := []int{100, 2, 3, 4}
	subarray := findUnsortedSubarray(nums)
	fmt.Println(subarray)
}

func Test_Case4(t *testing.T) {
	nums := []int{1, 3, 2, 4, 5}
	subarray := findUnsortedSubarray(nums)
	fmt.Println(subarray)
}
