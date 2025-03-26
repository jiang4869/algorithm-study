package main

func canBeValid(s string, locked string) bool {
	size := len(s)
	if size%2 != 0 {
		return false
	}
	ans := true
	left := make([]int, 0)
	a := make([]int, 0)
	for i, c := range s {
		if locked[i] == '0' {
			a = append(a, i)
		} else {
			if c == '(' {
				left = append(left, i)
			} else if len(left) > 0 {
				left = left[:len(left)-1]
			} else {
				if len(a) > 0 {
					a = a[1:]
				} else {
					ans = false
					break
				}
			}
		}
	}

	if len(left) == 0 && len(a)%2 != 0 {
		ans = false
	} else {
		i := len(a) - 1
		j := len(left) - 1
		for j >= 0 && i >= 0 {
			if a[i] < left[j] {
				ans = false
				break
			}
			i--
			j--
		}
		if ans {
			ans = (len(a)-len(left))%2 == 0 && j == -1
		}
	}

	return ans
}
