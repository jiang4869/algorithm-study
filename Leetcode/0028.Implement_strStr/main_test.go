package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	index := strStr(haystack, needle)
	fmt.Println(index)
}
