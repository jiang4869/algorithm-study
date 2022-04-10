package main

import (
	"fmt"
	"testing"
)

func Test_Case1(t *testing.T) {
	current := "02:30"
	correct := "04:35"
	time := convertTime(current, correct)
	fmt.Println(time)
}

func Test_Case2(t *testing.T) {
	current := "11:00"
	correct := "11:01"
	time := convertTime(current, correct)
	fmt.Println(time)
}
