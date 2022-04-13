package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	n := 3
	parenthesis := generateParenthesis(n)
	for _, val := range parenthesis {
		fmt.Println(val)
	}
}
