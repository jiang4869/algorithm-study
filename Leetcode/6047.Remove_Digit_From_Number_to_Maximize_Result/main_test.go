package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	number := "123"
	digit := '3'
	s := removeDigit(number, byte(digit))
	fmt.Println(s)
}
func Test_Case2(t *testing.T) {
	number := "1231"
	digit := '1'
	s := removeDigit(number, byte(digit))
	fmt.Println(s)
}

func Test_Case3(t *testing.T) {
	number := "551"
	digit := '5'
	s := removeDigit(number, byte(digit))
	fmt.Println(s)
}
