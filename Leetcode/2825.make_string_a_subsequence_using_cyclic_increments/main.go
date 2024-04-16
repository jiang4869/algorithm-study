package main

func canMakeSubsequence(str1 string, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	}
	j := 0
	for _, ch := range str1 {
		up := byte(ch) + 1
		if ch == 'z' {
			up = 'a'
		}
		if up == str2[j] || byte(ch) == str2[j] {
			j++
			if j == len(str2) {
				return true
			}
		}
	}

	return false
}
