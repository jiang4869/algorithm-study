package main

import "testing"

func TestCase1(t *testing.T) {
	n := 5
	f := [][]int{{0, 0, 1}, {0, 2, 2}, {1, 3, 2}}
	maximizeTheProfit(n, f)
}
