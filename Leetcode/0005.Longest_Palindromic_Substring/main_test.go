package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	str := "babad"
	ans := longestPalindrome(str)
	fmt.Println(ans)
}

func Test_Case2(t *testing.T) {
	str := "aa"
	ans := longestPalindrome(str)
	fmt.Println(ans)
}

func Test_Case3(t *testing.T) {
	str := "aaa"
	ans := longestPalindrome(str)
	fmt.Println(ans)
}
