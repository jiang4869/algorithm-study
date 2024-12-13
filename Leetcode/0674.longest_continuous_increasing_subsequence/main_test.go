package main

import "testing"

func TestCase1(t *testing.T) {
	arr := []int{1, 3, 5, 4, 7}
	ans := findLengthOfLCIS(arr)
	if ans != 3 {
		t.Errorf("ans %d not equal 3", ans)
	}
}

func TestCase2(t *testing.T) {
	arr := []int{1, 3, 5, 4, 2, 3, 4, 5}
	ans := findLengthOfLCIS(arr)
	if ans != 3 {
		t.Errorf("ans %d not equal 3", ans)
	}
}
