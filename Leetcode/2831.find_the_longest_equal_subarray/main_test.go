package main

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 3, 2, 3, 1, 3}
	k := 3
	longestEqualSubarray(nums, k)
}

func TestCase2(t *testing.T) {
	nums := []int{1}
	k := 0
	fmt.Println(longestEqualSubarray(nums, k))
}

func TestCase3(t *testing.T) {
	nums := []int{6, 4, 7, 6, 4, 8, 6, 6}
	k := 1
	fmt.Println(longestEqualSubarray(nums, k))
}

func TestCase4(t *testing.T) {
	nums := []int{9, 5, 1, 4, 2, 3, 4, 13, 5, 2, 7, 2, 2, 14}
	k := 4
	fmt.Println(longestEqualSubarray(nums, k))
}
