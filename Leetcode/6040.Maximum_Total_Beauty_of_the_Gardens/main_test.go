package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	newFlowers := 7
	target := 6
	full := 12
	partial := 1
	flowers := []int{1, 3, 1, 1}
	beauty := maximumBeauty(flowers, int64(newFlowers), target, full, partial)
	fmt.Println(beauty)
}

func Test_Case2(t *testing.T) {
	newFlowers := 10
	target := 5
	full := 2
	partial := 6
	flowers := []int{2, 4, 5, 3}
	beauty := maximumBeauty(flowers, int64(newFlowers), target, full, partial)
	fmt.Println(beauty)
}

func Test_Case3(t *testing.T) {
	newFlowers := 2
	target := 20
	full := 10
	partial := 2
	flowers := []int{20, 1, 15, 17, 10, 2, 4, 16, 15, 11}
	beauty := maximumBeauty(flowers, int64(newFlowers), target, full, partial)
	fmt.Println(beauty)
}
