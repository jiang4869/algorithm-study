package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	s := "abbca"
	sum := appealSum(s)
	fmt.Println(sum)
}
