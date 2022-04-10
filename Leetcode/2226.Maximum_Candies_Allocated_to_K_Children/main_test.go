package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	candies := []int{5, 8, 6}
	var k int64
	k = 3
	res := maximumCandies(candies, k)
	fmt.Println(res)
}

func Test_Case2(t *testing.T) {
	candies := []int{2, 5}
	var k int64
	k = 11
	res := maximumCandies(candies, k)
	fmt.Println(res)
}

func Test_Case3(t *testing.T) {
	candies := []int{1, 2, 3, 4, 10}
	var k int64
	k = 5
	res := maximumCandies(candies, k)
	fmt.Println(res)
}
