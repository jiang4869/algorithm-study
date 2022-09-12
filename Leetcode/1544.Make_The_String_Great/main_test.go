package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	s := "leEeetcode"
	good := makeGood(s)
	fmt.Println(good)
}

func Test_Case2(t *testing.T) {
	s := "abBAcC"
	good := makeGood(s)
	fmt.Println(good)
}
