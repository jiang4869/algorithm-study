package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 6}
	key := 7
	idx := LowerBound(nums, 0, len(nums), key)
	fmt.Println(idx)
}
