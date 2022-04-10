package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans := maxArea(height)
	fmt.Println(ans)
}

func Test_Case2(t *testing.T) {
	height := []int{1, 1}
	ans := maxArea(height)
	fmt.Println(ans)
}
