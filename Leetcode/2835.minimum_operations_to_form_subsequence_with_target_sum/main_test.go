package main

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 8}
	target := 7
	fmt.Println(minOperations(nums, target))
}

func TestCase2(t *testing.T) {
	nums := []int{16, 64, 4, 128}
	target := 6
	fmt.Println(minOperations(nums, target))
}
