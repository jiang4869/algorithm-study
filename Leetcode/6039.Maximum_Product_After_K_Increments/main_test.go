package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	nums := []int{0, 4}
	k := 5
	product := maximumProduct(nums, k)
	fmt.Println(product)
}

func Test_Case2(t *testing.T) {
	nums := []int{6, 3, 3, 2}
	k := 2
	product := maximumProduct(nums, k)
	fmt.Println(product)
}
