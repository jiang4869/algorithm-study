package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	cards := []int{3, 4, 2, 3, 4, 7}
	res := minimumCardPickup(cards)
	fmt.Println(res)
}
func Test_Case2(t *testing.T) {
	cards := []int{1, 0, 5, 3}
	res := minimumCardPickup(cards)
	fmt.Println(res)
}

func Test_Case3(t *testing.T) {
	cards := []int{95, 11, 8, 65, 5, 86, 30, 27, 30, 73, 15, 91, 30, 7, 37, 26, 55, 76, 60, 43, 36, 85, 47, 96, 6}
	res := minimumCardPickup(cards)
	fmt.Println(res)
}
