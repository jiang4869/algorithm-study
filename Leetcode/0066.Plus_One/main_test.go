package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	digits := []int{1, 2, 3}
	res := plusOne(digits)
	fmt.Println(res)
}
func Test_Case2(t *testing.T) {
	digits := []int{9}
	res := plusOne(digits)
	fmt.Println(res)
}

func Test_Case3(t *testing.T) {
	digits := []int{4, 3, 2, 1}
	res := plusOne(digits)
	fmt.Println(res)
}
