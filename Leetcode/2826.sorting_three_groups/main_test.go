package main

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{2, 2, 2, 2, 3, 3}
	fmt.Println(minimumOperations(nums))
}

func TestCase2(t *testing.T) {
	nums := []int{2, 1, 3, 2, 1}
	fmt.Println(minimumOperations(nums))
}
