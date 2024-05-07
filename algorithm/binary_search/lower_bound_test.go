package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 6}
	key := 4
	idx := LowerBound(nums, 0, len(nums), key)
	fmt.Println(idx)
}

func Test_Case2(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 4, 4, 6, 7, 8, 9}
	key := 4
	idx := LowerBoundBySort(nums, key)
	fmt.Println(idx)
}

func Test_Case3(t *testing.T) {
	nums := []int{1, 2, 3, 3, 4, 4, 4, 6, 7, 8, 9}
	key := 4
	idx := UpperBoundBySort(nums, key)
	fmt.Println(idx)
}
