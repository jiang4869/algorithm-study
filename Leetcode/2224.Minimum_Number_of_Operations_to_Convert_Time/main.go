package main

import (
	"strconv"
	"strings"
)

func convertTime(current string, correct string) int {
	s1 := strings.Split(current, ":")
	s2 := strings.Split(correct, ":")

	h1, _ := strconv.Atoi(s1[0])
	m1, _ := strconv.Atoi(s1[1])

	h2, _ := strconv.Atoi(s2[0])
	m2, _ := strconv.Atoi(s2[1])

	det := (h2*60 + m2) - (h1*60 + m1)

	cnt := 0
	for det > 0 {
		if det >= 60 {
			det -= 60
		} else if det >= 15 {
			det -= 15
		} else if det >= 5 {
			det -= 5
		} else {
			det -= 1
		}
		cnt++
	}
	return cnt
}
