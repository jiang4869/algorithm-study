package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	expression := "247+38"
	result := minimizeResult(expression)
	fmt.Println(result)
}

func Test_Case2(t *testing.T) {
	expression := "999+999"
	result := minimizeResult(expression)
	fmt.Println(result)
}

func Test_Case3(t *testing.T) {
	expression := "12+34"
	result := minimizeResult(expression)
	fmt.Println(result)
}
