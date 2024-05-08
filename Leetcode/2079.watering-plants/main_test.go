package main

import "testing"

func Test_Case1(t *testing.T) {
	plants := []int{2, 2, 3, 3}
	capacity := 5
	ans := wateringPlants(plants, capacity)
	if ans != 14 {
		t.Errorf("error,ans is %d not 14", ans)
	}
}
