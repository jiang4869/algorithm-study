package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	nums := []int{2, 3, 3, 2, 2}
	k := 2
	p := 2
	distinct := countDistinct(nums, k, p)
	fmt.Println(distinct)
}
func Test_Case2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	k := 4
	p := 1
	distinct := countDistinct(nums, k, p)
	fmt.Println(distinct)
}
func Test_Case3(t *testing.T) {
	nums := []int{1, 9, 8, 7, 19}
	k := 1
	p := 6
	distinct := countDistinct(nums, k, p)
	fmt.Println(distinct)

}
