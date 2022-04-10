package main

func longestPalindrome(s string) string {
	ans := s[:1]
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			length := check(i-1, i+1, s)
			if length+1 > len(ans) {
				ans = s[i-(length/2) : i+(length/2)+1]
			}
			length = check(i, i+1, s)
			if length > len(ans) {
				ans = s[i-(length/2)+1 : i+(length/2)+1]
			}
		}
	}
	return ans
}

func check(l, r int, str string) int {
	length := 0
	for l >= 0 && r < len(str) {
		if str[l] != str[r] {
			break
		}
		length += 2
		l--
		r++
	}
	return length
}
