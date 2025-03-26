package main

func minSwaps(s string) int {
	count := 0
	for _, char := range s {
		if char == '[' {
			count++
		} else if char == ']' && count > 0 {
			count--
		} else {
			count++
		}
	}

	return count / 2
}
