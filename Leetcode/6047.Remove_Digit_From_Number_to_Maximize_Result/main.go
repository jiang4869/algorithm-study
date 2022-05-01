package main

func removeDigit(number string, digit byte) string {
	res := ""
	for idx := range number {
		if number[idx] == digit {
			tmp := number[:idx] + number[idx+1:]
			if len(tmp) > len(res) {
				res = tmp
			} else if len(tmp) == len(res) {
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}
