package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3

	ans := convert(s, numRows)
	fmt.Println(ans)
}

func Test_Case2(t *testing.T) {
	s := "abababab"
	numRows := 2

	ans := convert(s, numRows)
	fmt.Println(ans)
}
