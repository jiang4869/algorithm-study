package main

func plusOne(digits []int) []int {
	i := len(digits) - 1
	digits[i] = digits[i] + 1
	for i >= 0 {
		tmp := digits[i] / 10
		digits[i] %= 10
		if tmp == 0 {
			break
		}
		if i == 0 {
			digits = append([]int{tmp}, digits...)
		} else {
			digits[i-1] += tmp
		}

		i--
	}
	return digits
}
