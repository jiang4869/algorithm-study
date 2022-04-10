package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	num := 1234
	integer := largestInteger(num)
	fmt.Println(integer)
}

func Test_Case2(t *testing.T) {
	num := 65875
	integer := largestInteger(num)
	fmt.Println(integer)
}
