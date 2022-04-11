package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	n := 2
	res := countNumbersWithUniqueDigits(n)
	fmt.Println(res)
}

func Test_Case2(t *testing.T) {
	n := 0
	res := countNumbersWithUniqueDigits(n)
	fmt.Println(res)
}

func Test_Case3(t *testing.T) {
	n := 8
	res := countNumbersWithUniqueDigits(n)
	fmt.Println(res)
}
