package main

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func furthestDistanceFromOrigin(moves string) int {
	l := 0
	r := 0
	t := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'L' {
			l++
		} else if moves[i] == 'R' {
			r++
		} else {
			t++
		}
	}
	return abs(l-r) + t
}
