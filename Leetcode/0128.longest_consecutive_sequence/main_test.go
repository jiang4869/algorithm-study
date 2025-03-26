package main

import "testing"

func TestCase1(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	ans := longestConsecutive(nums)
	if ans != 4 {
		t.Logf("require ans is 4, but got %d", ans)
	}
}
